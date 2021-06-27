package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/gorilla/mux"

	"github.com/Oguzyildirim/go-crud/internal"
	"github.com/Oguzyildirim/go-crud/internal/rest"
	"github.com/Oguzyildirim/go-crud/internal/rest/resttesting"
)

func TestUsers_Delete(t *testing.T) {
	t.Parallel()

	type output struct {
		expectedStatus int
		expected       interface{}
		target         interface{}
	}

	tests := []struct {
		name   string
		setup  func(*resttesting.FakeUserService)
		output output
	}{
		{
			"OK: 200",
			func(s *resttesting.FakeUserService) {},
			output{
				http.StatusOK,
				&struct{}{},
				&struct{}{},
			},
		},
		{
			"ERR: 404",
			func(s *resttesting.FakeUserService) {
				s.DeleteReturns(internal.NewErrorf(internal.ErrorCodeNotFound, "not found"))
			},
			output{
				http.StatusNotFound,
				&struct{}{},
				&struct{}{},
			},
		},
		{
			"ERR: 500",
			func(s *resttesting.FakeUserService) {
				s.DeleteReturns(errors.New("service failed"))
			},
			output{
				http.StatusInternalServerError,
				&struct{}{},
				&struct{}{},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			router := mux.NewRouter()
			svc := &resttesting.FakeUserService{}
			tt.setup(svc)

			rest.NewUserHandler(svc).Register(router)

			res := doRequest(router,
				httptest.NewRequest(http.MethodDelete, "/users/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", nil))

			assertResponse(t, res, test{tt.output.expected, tt.output.target})

			if tt.output.expectedStatus != res.StatusCode {
				t.Fatalf("expected code %d, actual %d", tt.output.expectedStatus, res.StatusCode)
			}
		})
	}
}

func TestUsers_Post(t *testing.T) {
	t.Parallel()

	type output struct {
		expectedStatus int
		expected       interface{}
		target         interface{}
	}

	tests := []struct {
		name   string
		setup  func(*resttesting.FakeUserService)
		input  []byte
		output output
	}{
		{
			"OK: 201",
			func(s *resttesting.FakeUserService) {
				s.CreateReturns(
					internal.User{
						ID:       "1-2-3",
						Name:     "name",
						Lastname: "lastname",
						Username: "username",
						Country:  "UK",
					},
					nil)
			},
			func() []byte {
				b, _ := json.Marshal(&rest.CreateUsersRequest{
					Name:     "name",
					Lastname: "lastname",
					Username: "username",
					Country:  "UK",
				})

				return b
			}(),
			output{
				http.StatusCreated,
				&rest.CreateUsersResponse{
					User: rest.User{
						ID:       "1-2-3",
						Name:     "name",
						Lastname: "lastname",
						Username: "username",
						Country:  "UK",
					},
				},
				&rest.CreateUsersResponse{},
			},
		},
		{
			"ERR: 400",
			func(*resttesting.FakeUserService) {},
			[]byte(`{"invalid":"json`),
			output{
				http.StatusBadRequest,
				&rest.ErrorResponse{
					Error: "invalid request",
				},
				&rest.ErrorResponse{},
			},
		},
		{
			"ERR: 500",
			func(s *resttesting.FakeUserService) {
				s.CreateReturns(internal.User{},
					errors.New("service error"))
			},
			[]byte(`{}`),
			output{
				http.StatusInternalServerError,
				&rest.ErrorResponse{
					Error: "internal error",
				},
				&rest.ErrorResponse{},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			router := mux.NewRouter()
			svc := &resttesting.FakeUserService{}
			tt.setup(svc)

			rest.NewUserHandler(svc).Register(router)

			res := doRequest(router,
				httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(tt.input)))

			assertResponse(t, res, test{tt.output.expected, tt.output.target})

			if tt.output.expectedStatus != res.StatusCode {
				t.Fatalf("expected code %d, actual %d", tt.output.expectedStatus, res.StatusCode)
			}
		})
	}
}

func TestUsers_Read(t *testing.T) {
	t.Parallel()

	type output struct {
		expectedStatus int
		expected       interface{}
		target         interface{}
	}

	tests := []struct {
		name   string
		setup  func(*resttesting.FakeUserService)
		output output
	}{
		{
			"OK: 200",
			func(s *resttesting.FakeUserService) {
				s.FindReturns(
					internal.User{
						ID:       "a-b-c",
						Name:     "name",
						Lastname: "lastname",
						Username: "username",
						Country:  "UK",
					},
					nil)
			},
			output{
				http.StatusOK,
				&rest.ReadUserResponse{
					User: rest.User{
						ID:       "a-b-c",
						Name:     "name",
						Lastname: "lastname",
						Username: "username",
						Country:  "UK",
					},
				},
				&rest.ReadUserResponse{},
			},
		},
		{
			"OK: 200",
			func(s *resttesting.FakeUserService) {
				s.FindReturns(internal.User{},
					internal.NewErrorf(internal.ErrorCodeNotFound, "not found"))
			},
			output{
				http.StatusNotFound,
				&rest.ErrorResponse{
					Error: "find failed",
				},
				&rest.ErrorResponse{},
			},
		},
		{
			"ERR: 500",
			func(s *resttesting.FakeUserService) {
				s.FindReturns(internal.User{},
					errors.New("service error"))
			},
			output{
				http.StatusInternalServerError,
				&rest.ErrorResponse{
					Error: "internal error",
				},
				&rest.ErrorResponse{},
			},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			router := mux.NewRouter()
			svc := &resttesting.FakeUserService{}
			tt.setup(svc)

			rest.NewUserHandler(svc).Register(router)

			res := doRequest(router,
				httptest.NewRequest(http.MethodGet, "/users/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", nil))

			assertResponse(t, res, test{tt.output.expected, tt.output.target})

			if tt.output.expectedStatus != res.StatusCode {
				t.Fatalf("expected code %d, actual %d", tt.output.expectedStatus, res.StatusCode)
			}
		})
	}
}

func TestUsers_Update(t *testing.T) {
	t.Parallel()

	type output struct {
		expectedStatus int
		expected       interface{}
		target         interface{}
	}

	tests := []struct {
		name   string
		setup  func(*resttesting.FakeUserService)
		input  []byte
		output output
	}{
		{
			"OK: 200",
			func(s *resttesting.FakeUserService) {},
			func() []byte {
				b, _ := json.Marshal(&rest.UpdateUsersRequest{
					Name:     "name",
					Lastname: "lastname",
					Username: "username",
					Country:  "UK",
				})

				return b
			}(),
			output{
				http.StatusOK,
				&struct{}{},
				&struct{}{},
			},
		},
		{
			"ERR: 400",
			func(*resttesting.FakeUserService) {},
			[]byte(`{"invalid":"json`),
			output{
				http.StatusBadRequest,
				&rest.ErrorResponse{
					Error: "invalid request",
				},
				&rest.ErrorResponse{},
			},
		},
		{
			"ERR: 404",
			func(s *resttesting.FakeUserService) {
				s.UpdateReturns(internal.NewErrorf(internal.ErrorCodeNotFound, "not found"))
			},
			func() []byte {
				b, _ := json.Marshal(&rest.UpdateUsersRequest{
					Name:     "name",
					Lastname: "lastname",
					Username: "username",
					Country:  "UK",
				})

				return b
			}(),
			output{
				http.StatusNotFound,
				&struct{}{},
				&struct{}{},
			},
		},
		{
			"ERR: 500",
			func(s *resttesting.FakeUserService) {
				s.UpdateReturns(errors.New("service error"))
			},
			[]byte(`{}`),
			output{
				http.StatusInternalServerError,
				&rest.ErrorResponse{
					Error: "internal error",
				},
				&rest.ErrorResponse{},
			},
		},
	}

	//-

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			router := mux.NewRouter()
			svc := &resttesting.FakeUserService{}
			tt.setup(svc)

			rest.NewUserHandler(svc).Register(router)

			res := doRequest(router,
				httptest.NewRequest(http.MethodPut, "/users/aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee", bytes.NewReader(tt.input)))

			assertResponse(t, res, test{tt.output.expected, tt.output.target})

			if tt.output.expectedStatus != res.StatusCode {
				t.Fatalf("expected code %d, actual %d", tt.output.expectedStatus, res.StatusCode)
			}
		})
	}
}

type test struct {
	expected interface{}
	target   interface{}
}

func doRequest(router *mux.Router, req *http.Request) *http.Response {
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	return rr.Result()
}

func assertResponse(t *testing.T, res *http.Response, test test) {
	t.Helper()

	if err := json.NewDecoder(res.Body).Decode(test.target); err != nil {
		t.Fatalf("couldn't decode %s", err)
	}
	defer res.Body.Close()

	if !cmp.Equal(test.expected, test.target, cmpopts.IgnoreUnexported()) {
		t.Fatalf("expected results don't match: %s", cmp.Diff(test.expected, test.target, cmpopts.IgnoreUnexported()))
	}
}
