// Package openapi3 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package openapi3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// CreateUser request  with any body
	CreateUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateUser(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ReadUserByCountry request
	ReadUserByCountry(ctx context.Context, country string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteUser request
	DeleteUser(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ReadUser request
	ReadUser(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// UpdateUser request  with any body
	UpdateUserWithBody(ctx context.Context, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	UpdateUser(ctx context.Context, userId string, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) CreateUserWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateUserRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateUser(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateUserRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ReadUserByCountry(ctx context.Context, country string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewReadUserByCountryRequest(c.Server, country)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteUser(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteUserRequest(c.Server, userId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ReadUser(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewReadUserRequest(c.Server, userId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateUserWithBody(ctx context.Context, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateUserRequestWithBody(c.Server, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) UpdateUser(ctx context.Context, userId string, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewUpdateUserRequest(c.Server, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewCreateUserRequest calls the generic CreateUser builder with application/json body
func NewCreateUserRequest(server string, body CreateUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateUserRequestWithBody(server, "application/json", bodyReader)
}

// NewCreateUserRequestWithBody generates requests for CreateUser with any type of body
func NewCreateUserRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users")
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewReadUserByCountryRequest generates requests for ReadUserByCountry
func NewReadUserByCountryRequest(server string, country string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "country", runtime.ParamLocationPath, country)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/by-country/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDeleteUserRequest generates requests for DeleteUser
func NewDeleteUserRequest(server string, userId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewReadUserRequest generates requests for ReadUser
func NewReadUserRequest(server string, userId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewUpdateUserRequest calls the generic UpdateUser builder with application/json body
func NewUpdateUserRequest(server string, userId string, body UpdateUserJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewUpdateUserRequestWithBody(server, userId, "application/json", bodyReader)
}

// NewUpdateUserRequestWithBody generates requests for UpdateUser with any type of body
func NewUpdateUserRequestWithBody(server string, userId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "userId", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/users/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = operationPath[1:]
	}
	operationURL := url.URL{
		Path: operationPath,
	}

	queryURL := serverURL.ResolveReference(&operationURL)

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// CreateUser request  with any body
	CreateUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateUserResponse, error)

	CreateUserWithResponse(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateUserResponse, error)

	// ReadUserByCountry request
	ReadUserByCountryWithResponse(ctx context.Context, country string, reqEditors ...RequestEditorFn) (*ReadUserByCountryResponse, error)

	// DeleteUser request
	DeleteUserWithResponse(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*DeleteUserResponse, error)

	// ReadUser request
	ReadUserWithResponse(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*ReadUserResponse, error)

	// UpdateUser request  with any body
	UpdateUserWithBodyWithResponse(ctx context.Context, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error)

	UpdateUserWithResponse(ctx context.Context, userId string, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error)
}

type CreateUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *struct {
		User *User `json:"user,omitempty"`
	}
	JSON400 *struct {
		Error *string `json:"error,omitempty"`
	}
	JSON500 *struct {
		Error *string `json:"error,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r CreateUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ReadUserByCountryResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		User *User `json:"user,omitempty"`
	}
	JSON500 *struct {
		Error *string `json:"error,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ReadUserByCountryResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ReadUserByCountryResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON500      *struct {
		Error *string `json:"error,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r DeleteUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ReadUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		User *User `json:"user,omitempty"`
	}
	JSON500 *struct {
		Error *string `json:"error,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r ReadUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ReadUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type UpdateUserResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *struct {
		Error *string `json:"error,omitempty"`
	}
	JSON500 *struct {
		Error *string `json:"error,omitempty"`
	}
}

// Status returns HTTPResponse.Status
func (r UpdateUserResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r UpdateUserResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// CreateUserWithBodyWithResponse request with arbitrary body returning *CreateUserResponse
func (c *ClientWithResponses) CreateUserWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateUserResponse, error) {
	rsp, err := c.CreateUserWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateUserResponse(rsp)
}

func (c *ClientWithResponses) CreateUserWithResponse(ctx context.Context, body CreateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateUserResponse, error) {
	rsp, err := c.CreateUser(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateUserResponse(rsp)
}

// ReadUserByCountryWithResponse request returning *ReadUserByCountryResponse
func (c *ClientWithResponses) ReadUserByCountryWithResponse(ctx context.Context, country string, reqEditors ...RequestEditorFn) (*ReadUserByCountryResponse, error) {
	rsp, err := c.ReadUserByCountry(ctx, country, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseReadUserByCountryResponse(rsp)
}

// DeleteUserWithResponse request returning *DeleteUserResponse
func (c *ClientWithResponses) DeleteUserWithResponse(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*DeleteUserResponse, error) {
	rsp, err := c.DeleteUser(ctx, userId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteUserResponse(rsp)
}

// ReadUserWithResponse request returning *ReadUserResponse
func (c *ClientWithResponses) ReadUserWithResponse(ctx context.Context, userId string, reqEditors ...RequestEditorFn) (*ReadUserResponse, error) {
	rsp, err := c.ReadUser(ctx, userId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseReadUserResponse(rsp)
}

// UpdateUserWithBodyWithResponse request with arbitrary body returning *UpdateUserResponse
func (c *ClientWithResponses) UpdateUserWithBodyWithResponse(ctx context.Context, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error) {
	rsp, err := c.UpdateUserWithBody(ctx, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateUserResponse(rsp)
}

func (c *ClientWithResponses) UpdateUserWithResponse(ctx context.Context, userId string, body UpdateUserJSONRequestBody, reqEditors ...RequestEditorFn) (*UpdateUserResponse, error) {
	rsp, err := c.UpdateUser(ctx, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseUpdateUserResponse(rsp)
}

// ParseCreateUserResponse parses an HTTP response from a CreateUserWithResponse call
func ParseCreateUserResponse(rsp *http.Response) (*CreateUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &CreateUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest struct {
			User *User `json:"user,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest struct {
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseReadUserByCountryResponse parses an HTTP response from a ReadUserByCountryWithResponse call
func ParseReadUserByCountryResponse(rsp *http.Response) (*ReadUserByCountryResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ReadUserByCountryResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			User *User `json:"user,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseDeleteUserResponse parses an HTTP response from a DeleteUserWithResponse call
func ParseDeleteUserResponse(rsp *http.Response) (*DeleteUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &DeleteUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseReadUserResponse parses an HTTP response from a ReadUserWithResponse call
func ParseReadUserResponse(rsp *http.Response) (*ReadUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ReadUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			User *User `json:"user,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}

// ParseUpdateUserResponse parses an HTTP response from a UpdateUserWithResponse call
func ParseUpdateUserResponse(rsp *http.Response) (*UpdateUserResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &UpdateUserResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest struct {
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 500:
		var dest struct {
			Error *string `json:"error,omitempty"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON500 = &dest

	}

	return response, nil
}
