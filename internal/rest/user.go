package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Oguzyildirim/go-crud/internal"
)

const uuidRegEx string = `[0-9a-fA-F]{8}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{4}\-[0-9a-fA-F]{12}`

const countryRegEx string = `[a-zA-Z\\s\'\"]+`

//go:generate counterfeiter -o resttesting/user_service.gen.go . UserService

// UserService
type UserService interface {
	Create(ctx context.Context, name string, lastname string, username string, country string) (internal.User, error)
	Delete(ctx context.Context, id string) error
	Find(ctx context.Context, id string) (internal.User, error)
	FindByCountry(ctx context.Context, country string) ([]internal.User, error)
	Update(ctx context.Context, id string, name string, lastname string, username string, country string) error
}

// UserHandler
type UserHandler struct {
	svc UserService
}

// NewUserHandler
func NewUserHandler(svc UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

// Register connects the handlers to the router.
func (u *UserHandler) Register(r *mux.Router) {
	r.HandleFunc("/users", u.create).Methods(http.MethodPost)
	r.HandleFunc(fmt.Sprintf("/users/{id:%s}", uuidRegEx), u.find).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("/users/by-country/{country:%s}", countryRegEx), u.findByCountry).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("/users/{id:%s}", uuidRegEx), u.update).Methods(http.MethodPut)
	r.HandleFunc(fmt.Sprintf("/users/{id:%s}", uuidRegEx), u.delete).Methods(http.MethodDelete)
}

// User is an actor
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Country  string `json:"country"`
}

// CreateUsersRequest defines the request used for creating users.
type CreateUsersRequest struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Country  string `json:"country"`
}

// CreateUsersResponse defines the response returned back after creating users.
type CreateUsersResponse struct {
	User User `json:"user"`
}

func (u *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	var req CreateUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		renderErrorResponse(r.Context(), w, "invalid request", internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "json decoder"))
		return
	}

	defer r.Body.Close()

	user, err := u.svc.Create(r.Context(), req.Name, req.Lastname, req.Username, req.Country)
	fmt.Println(err)
	if err != nil {
		renderErrorResponse(r.Context(), w, "create failed", err)
		return
	}

	renderResponse(w,
		&CreateUsersResponse{
			User: User{
				ID:       user.ID,
				Name:     user.Name,
				Lastname: user.Lastname,
				Username: user.Username,
				Country:  user.Country,
			},
		},
		http.StatusCreated)
}

func (u *UserHandler) delete(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]

	if err := u.svc.Delete(r.Context(), id); err != nil {
		renderErrorResponse(r.Context(), w, "delete failed", err)
		return
	}

	renderResponse(w, struct{}{}, http.StatusOK)
}

// ReadUsersResponse defines the response returned back after searching one user.
type ReadUserResponse struct {
	User User `json:"user"`
}

func (u *UserHandler) find(w http.ResponseWriter, r *http.Request) {
	id, _ := mux.Vars(r)["id"]

	user, err := u.svc.Find(r.Context(), id)
	if err != nil {
		renderErrorResponse(r.Context(), w, "find failed", err)
		return
	}

	renderResponse(w,
		&ReadUserResponse{
			User: User{
				ID:       user.ID,
				Name:     user.Name,
				Lastname: user.Lastname,
				Username: user.Username,
				Country:  user.Country,
			},
		},
		http.StatusOK)
}

// ReadUsersResponse defines the response returned back after searching one user.
type ReadUsersResponse struct {
	User []User `json:"users"`
}

func (u *UserHandler) findByCountry(w http.ResponseWriter, r *http.Request) {
	fmt.Print("hit")
	country, _ := mux.Vars(r)["country"]

	users, err := u.svc.FindByCountry(r.Context(), country)
	fmt.Print(users)
	if err != nil {
		renderErrorResponse(r.Context(), w, "find failed", err)
		return
	}
	var usersResponse []User
	for _, value := range users {
		user := User{
			ID:       value.ID,
			Name:     value.Name,
			Lastname: value.Lastname,
			Username: value.Username,
			Country:  value.Country,
		}
		usersResponse = append(usersResponse, user)
	}
	renderResponse(w,
		&ReadUsersResponse{
			usersResponse,
		},
		http.StatusOK)
}

// UpdateUsersRequest defines the request used for updating a user.
type UpdateUsersRequest struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Country  string `json:"country"`
}

func (u *UserHandler) update(w http.ResponseWriter, r *http.Request) {
	var req UpdateUsersRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		renderErrorResponse(r.Context(), w, "invalid request", internal.WrapErrorf(err, internal.ErrorCodeInvalidArgument, "json decoder"))
		return
	}

	defer r.Body.Close()

	id, _ := mux.Vars(r)["id"]
	err := u.svc.Update(r.Context(), id, req.Name, req.Lastname, req.Username, req.Country)
	if err != nil {
		renderErrorResponse(r.Context(), w, "update failed", err)
		return
	}

	renderResponse(w, &struct{}{}, http.StatusOK)
}
