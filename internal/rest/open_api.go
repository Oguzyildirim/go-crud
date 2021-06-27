package rest

import (
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
	"github.com/gorilla/mux"
)

//go:generate go run ../../cmd/openapi-gen/main.go -path .
//go:generate oapi-codegen -package openapi3 -generate types  -o ../../pkg/openapi3/user_types.gen.go openapi3.yaml
//go:generate oapi-codegen -package openapi3 -generate client -o ../../pkg/openapi3/client.gen.go     openapi3.yaml

// NewOpenAPI3 instantiates the OpenAPI specification for this service.
func NewOpenAPI3() openapi3.T {
	swagger := openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "User API",
			Description: "REST APIs used for interacting with the User Service",
			Version:     "0.0.0",
			License: &openapi3.License{
				Name: "MIT",
				URL:  "https://opensource.org/licenses/MIT",
			},
			Contact: &openapi3.Contact{
				URL: "https://github.com/Oguzyildirim/go-crud",
			},
		},
		Servers: openapi3.Servers{
			&openapi3.Server{
				Description: "Local development",
				URL:         "http://127.0.0.1:9234",
			},
		},
	}

	swagger.Components.Schemas = openapi3.Schemas{
		"User": openapi3.NewSchemaRef("",
			openapi3.NewObjectSchema().
				WithProperty("id", openapi3.NewUUIDSchema()).
				WithProperty("name", openapi3.NewStringSchema()).
				WithProperty("lastname", openapi3.NewStringSchema()).
				WithProperty("username", openapi3.NewStringSchema()).
				WithProperty("country", openapi3.NewStringSchema())),
	}

	swagger.Components.RequestBodies = openapi3.RequestBodies{
		"CreateUsersRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for creating a user.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("name", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("lastname", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("username", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("country", openapi3.NewStringSchema().
						WithMinLength(1)),
				),
		},
		"UpdateUsersRequest": &openapi3.RequestBodyRef{
			Value: openapi3.NewRequestBody().
				WithDescription("Request used for updating a user.").
				WithRequired(true).
				WithJSONSchema(openapi3.NewSchema().
					WithProperty("name", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("lastname", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("username", openapi3.NewStringSchema().
						WithMinLength(1)).
					WithProperty("country", openapi3.NewStringSchema().
						WithMinLength(1)),
				),
		},
	}

	swagger.Components.Responses = openapi3.Responses{
		"ErrorResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response when errors happen.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithProperty("error", openapi3.NewStringSchema()))),
		},
		"CreateUsersResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after creating users.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("user", &openapi3.SchemaRef{
						Ref: "#/components/schemas/User",
					}))),
		},
		"ReadUsersResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after searching one user.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("user", &openapi3.SchemaRef{
						Ref: "#/components/schemas/User",
					}))),
		},
		"ReadUsersByCountryResponse": &openapi3.ResponseRef{
			Value: openapi3.NewResponse().
				WithDescription("Response returned back after searching users by country.").
				WithContent(openapi3.NewContentWithJSONSchema(openapi3.NewSchema().
					WithPropertyRef("user", &openapi3.SchemaRef{
						Ref: "#/components/schemas/User",
					}))),
		},
	}

	swagger.Paths = openapi3.Paths{
		"/users": &openapi3.PathItem{
			Post: &openapi3.Operation{
				OperationID: "CreateUser",
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/CreateUsersRequest",
				},
				Responses: openapi3.Responses{
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"201": &openapi3.ResponseRef{
						Ref: "#/components/responses/CreateUsersResponse",
					},
				},
			},
		},
		"/users/{userId}": &openapi3.PathItem{
			Delete: &openapi3.Operation{
				OperationID: "DeleteUser",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("userId").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("User updated"),
					},
					"404": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("User not found"),
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
			Get: &openapi3.Operation{
				OperationID: "ReadUser",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("userId").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/ReadUsersResponse",
					},
					"404": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("User not found"),
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
			Put: &openapi3.Operation{
				OperationID: "UpdateUser",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("userId").
							WithSchema(openapi3.NewUUIDSchema()),
					},
				},
				RequestBody: &openapi3.RequestBodyRef{
					Ref: "#/components/requestBodies/UpdateUsersRequest",
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("User updated"),
					},
					"400": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
					"404": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("User not found"),
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
		},
		"/users/by-country/{country}": &openapi3.PathItem{
			Get: &openapi3.Operation{
				OperationID: "ReadUserByCountry",
				Parameters: []*openapi3.ParameterRef{
					{
						Value: openapi3.NewPathParameter("country").
							WithSchema(openapi3.NewStringSchema()),
					},
				},
				Responses: openapi3.Responses{
					"200": &openapi3.ResponseRef{
						Ref: "#/components/responses/ReadUsersByCountryResponse",
					},
					"404": &openapi3.ResponseRef{
						Value: openapi3.NewResponse().WithDescription("Users not found"),
					},
					"500": &openapi3.ResponseRef{
						Ref: "#/components/responses/ErrorResponse",
					},
				},
			},
		},
	}

	return swagger
}

func RegisterOpenAPI(r *mux.Router) {
	swagger := NewOpenAPI3()

	r.HandleFunc("/openapi3.json", func(w http.ResponseWriter, r *http.Request) {
		renderResponse(w, &swagger, http.StatusOK)
	}).Methods(http.MethodGet)

	r.HandleFunc("/openapi3.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-yaml")

		data, _ := yaml.Marshal(&swagger)

		_, _ = w.Write(data)

		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodGet)
}
