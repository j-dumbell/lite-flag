// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package oapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// Defines values for ApiKeyRole.
const (
	ApiKeyRoleAdmin    ApiKeyRole = "admin"
	ApiKeyRoleReadonly ApiKeyRole = "readonly"
	ApiKeyRoleRoot     ApiKeyRole = "root"
)

// Defines values for ApiKeyInputRole.
const (
	ApiKeyInputRoleAdmin    ApiKeyInputRole = "admin"
	ApiKeyInputRoleReadonly ApiKeyInputRole = "readonly"
	ApiKeyInputRoleRoot     ApiKeyInputRole = "root"
)

// Defines values for FlagType.
const (
	FlagTypeBoolean FlagType = "boolean"
	FlagTypeString  FlagType = "string"
)

// Defines values for FlagInputType.
const (
	FlagInputTypeBoolean FlagInputType = "boolean"
	FlagInputTypeString  FlagInputType = "string"
)

// ApiKey defines model for ApiKey.
type ApiKey struct {
	// Key The API Key
	Key string `json:"key"`

	// Name Unique identifier for the API key
	Name string `json:"name"`

	// Role The associated permissions of the key
	Role ApiKeyRole `json:"role"`
}

// ApiKeyRole The associated permissions of the key
type ApiKeyRole string

// ApiKeyInput defines model for ApiKeyInput.
type ApiKeyInput struct {
	// Name Unique identifier for the API key
	Name string `json:"name"`

	// Role The associated permissions of the key
	Role ApiKeyInputRole `json:"role"`
}

// ApiKeyInputRole The associated permissions of the key
type ApiKeyInputRole string

// Flag defines model for Flag.
type Flag struct {
	// IsPublic Whether or not the flag is public.
	IsPublic bool `json:"isPublic"`

	// Key Unique identifier for the feature flag
	Key string `json:"key"`

	// Type Type of the flag value
	Type FlagType `json:"type"`

	// Value Value of the flag, must match the type
	Value Flag_Value `json:"value"`
}

// FlagType Type of the flag value
type FlagType string

// FlagValue0 defines model for .
type FlagValue0 = bool

// FlagValue1 defines model for .
type FlagValue1 = string

// Flag_Value Value of the flag, must match the type
type Flag_Value struct {
	union json.RawMessage
}

// FlagInput defines model for FlagInput.
type FlagInput struct {
	// IsPublic Whether the feature flag is public or not
	IsPublic bool `json:"isPublic"`

	// Type Type of the flag value
	Type FlagInputType `json:"type"`

	// Value Value of the flag, must match the type
	Value FlagInput_Value `json:"value"`
}

// FlagInputType Type of the flag value
type FlagInputType string

// FlagInputValue0 defines model for .
type FlagInputValue0 = bool

// FlagInputValue1 defines model for .
type FlagInputValue1 = string

// FlagInput_Value Value of the flag, must match the type
type FlagInput_Value struct {
	union json.RawMessage
}

// HealthResponse defines model for HealthResponse.
type HealthResponse struct {
	// Database Whether the database is healthy
	Database bool `json:"database"`
}

// PostApiKeysJSONRequestBody defines body for PostApiKeys for application/json ContentType.
type PostApiKeysJSONRequestBody = ApiKeyInput

// PostFlagsJSONRequestBody defines body for PostFlags for application/json ContentType.
type PostFlagsJSONRequestBody = Flag

// PutFlagsKeyJSONRequestBody defines body for PutFlagsKey for application/json ContentType.
type PutFlagsKeyJSONRequestBody = FlagInput

// AsFlagValue0 returns the union data inside the Flag_Value as a FlagValue0
func (t Flag_Value) AsFlagValue0() (FlagValue0, error) {
	var body FlagValue0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFlagValue0 overwrites any union data inside the Flag_Value as the provided FlagValue0
func (t *Flag_Value) FromFlagValue0(v FlagValue0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFlagValue0 performs a merge with any union data inside the Flag_Value, using the provided FlagValue0
func (t *Flag_Value) MergeFlagValue0(v FlagValue0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFlagValue1 returns the union data inside the Flag_Value as a FlagValue1
func (t Flag_Value) AsFlagValue1() (FlagValue1, error) {
	var body FlagValue1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFlagValue1 overwrites any union data inside the Flag_Value as the provided FlagValue1
func (t *Flag_Value) FromFlagValue1(v FlagValue1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFlagValue1 performs a merge with any union data inside the Flag_Value, using the provided FlagValue1
func (t *Flag_Value) MergeFlagValue1(v FlagValue1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t Flag_Value) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Flag_Value) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFlagInputValue0 returns the union data inside the FlagInput_Value as a FlagInputValue0
func (t FlagInput_Value) AsFlagInputValue0() (FlagInputValue0, error) {
	var body FlagInputValue0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFlagInputValue0 overwrites any union data inside the FlagInput_Value as the provided FlagInputValue0
func (t *FlagInput_Value) FromFlagInputValue0(v FlagInputValue0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFlagInputValue0 performs a merge with any union data inside the FlagInput_Value, using the provided FlagInputValue0
func (t *FlagInput_Value) MergeFlagInputValue0(v FlagInputValue0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsFlagInputValue1 returns the union data inside the FlagInput_Value as a FlagInputValue1
func (t FlagInput_Value) AsFlagInputValue1() (FlagInputValue1, error) {
	var body FlagInputValue1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFlagInputValue1 overwrites any union data inside the FlagInput_Value as the provided FlagInputValue1
func (t *FlagInput_Value) FromFlagInputValue1(v FlagInputValue1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFlagInputValue1 performs a merge with any union data inside the FlagInput_Value, using the provided FlagInputValue1
func (t *FlagInput_Value) MergeFlagInputValue1(v FlagInputValue1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t FlagInput_Value) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FlagInput_Value) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new API key
	// (POST /api-keys)
	PostApiKeys(w http.ResponseWriter, r *http.Request)
	// Rotate an API key
	// (POST /api-keys/{name}/rotate)
	PostApiKeysNameRotate(w http.ResponseWriter, r *http.Request, name string)
	// Retrieve all feature flags
	// (GET /flags)
	GetFlags(w http.ResponseWriter, r *http.Request)
	// Create a new feature flag
	// (POST /flags)
	PostFlags(w http.ResponseWriter, r *http.Request)
	// Delete a feature flag
	// (DELETE /flags/{key})
	DeleteFlagsKey(w http.ResponseWriter, r *http.Request, key string)
	// Retrieve a single feature flag by key
	// (GET /flags/{key})
	GetFlagsKey(w http.ResponseWriter, r *http.Request, key string)
	// Update an existing feature flag
	// (PUT /flags/{key})
	PutFlagsKey(w http.ResponseWriter, r *http.Request, key string)
	// API healthcheck
	// (GET /healthz)
	GetHealthz(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Create a new API key
// (POST /api-keys)
func (_ Unimplemented) PostApiKeys(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Rotate an API key
// (POST /api-keys/{name}/rotate)
func (_ Unimplemented) PostApiKeysNameRotate(w http.ResponseWriter, r *http.Request, name string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Retrieve all feature flags
// (GET /flags)
func (_ Unimplemented) GetFlags(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a new feature flag
// (POST /flags)
func (_ Unimplemented) PostFlags(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete a feature flag
// (DELETE /flags/{key})
func (_ Unimplemented) DeleteFlagsKey(w http.ResponseWriter, r *http.Request, key string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Retrieve a single feature flag by key
// (GET /flags/{key})
func (_ Unimplemented) GetFlagsKey(w http.ResponseWriter, r *http.Request, key string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update an existing feature flag
// (PUT /flags/{key})
func (_ Unimplemented) PutFlagsKey(w http.ResponseWriter, r *http.Request, key string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// API healthcheck
// (GET /healthz)
func (_ Unimplemented) GetHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostApiKeys operation middleware
func (siw *ServerInterfaceWrapper) PostApiKeys(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, ApiKeyAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiKeys(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostApiKeysNameRotate operation middleware
func (siw *ServerInterfaceWrapper) PostApiKeysNameRotate(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "name" -------------
	var name string

	err = runtime.BindStyledParameterWithOptions("simple", "name", chi.URLParam(r, "name"), &name, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "name", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, ApiKeyAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostApiKeysNameRotate(w, r, name)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetFlags operation middleware
func (siw *ServerInterfaceWrapper) GetFlags(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, ApiKeyAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFlags(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PostFlags operation middleware
func (siw *ServerInterfaceWrapper) PostFlags(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, ApiKeyAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostFlags(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// DeleteFlagsKey operation middleware
func (siw *ServerInterfaceWrapper) DeleteFlagsKey(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "key" -------------
	var key string

	err = runtime.BindStyledParameterWithOptions("simple", "key", chi.URLParam(r, "key"), &key, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, ApiKeyAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteFlagsKey(w, r, key)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetFlagsKey operation middleware
func (siw *ServerInterfaceWrapper) GetFlagsKey(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "key" -------------
	var key string

	err = runtime.BindStyledParameterWithOptions("simple", "key", chi.URLParam(r, "key"), &key, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, ApiKeyAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetFlagsKey(w, r, key)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// PutFlagsKey operation middleware
func (siw *ServerInterfaceWrapper) PutFlagsKey(w http.ResponseWriter, r *http.Request) {

	var err error

	// ------------- Path parameter "key" -------------
	var key string

	err = runtime.BindStyledParameterWithOptions("simple", "key", chi.URLParam(r, "key"), &key, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "key", Err: err})
		return
	}

	ctx := r.Context()

	ctx = context.WithValue(ctx, ApiKeyAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutFlagsKey(w, r, key)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

// GetHealthz operation middleware
func (siw *ServerInterfaceWrapper) GetHealthz(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	ctx = context.WithValue(ctx, ApiKeyAuthScopes, []string{})

	r = r.WithContext(ctx)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetHealthz(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r)
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api-keys", wrapper.PostApiKeys)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/api-keys/{name}/rotate", wrapper.PostApiKeysNameRotate)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/flags", wrapper.GetFlags)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/flags", wrapper.PostFlags)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/flags/{key}", wrapper.DeleteFlagsKey)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/flags/{key}", wrapper.GetFlagsKey)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/flags/{key}", wrapper.PutFlagsKey)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/healthz", wrapper.GetHealthz)
	})

	return r
}

type PostApiKeysRequestObject struct {
	Body *PostApiKeysJSONRequestBody
}

type PostApiKeysResponseObject interface {
	VisitPostApiKeysResponse(w http.ResponseWriter) error
}

type PostApiKeys201JSONResponse ApiKey

func (response PostApiKeys201JSONResponse) VisitPostApiKeysResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostApiKeys400JSONResponse map[string]interface{}

func (response PostApiKeys400JSONResponse) VisitPostApiKeysResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostApiKeys401Response struct {
}

func (response PostApiKeys401Response) VisitPostApiKeysResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type PostApiKeys403JSONResponse map[string]interface{}

func (response PostApiKeys403JSONResponse) VisitPostApiKeysResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type PostApiKeys409Response struct {
}

func (response PostApiKeys409Response) VisitPostApiKeysResponse(w http.ResponseWriter) error {
	w.WriteHeader(409)
	return nil
}

type PostApiKeysNameRotateRequestObject struct {
	Name string `json:"name"`
}

type PostApiKeysNameRotateResponseObject interface {
	VisitPostApiKeysNameRotateResponse(w http.ResponseWriter) error
}

type PostApiKeysNameRotate201JSONResponse ApiKey

func (response PostApiKeysNameRotate201JSONResponse) VisitPostApiKeysNameRotateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostApiKeysNameRotate401Response struct {
}

func (response PostApiKeysNameRotate401Response) VisitPostApiKeysNameRotateResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type PostApiKeysNameRotate403JSONResponse map[string]interface{}

func (response PostApiKeysNameRotate403JSONResponse) VisitPostApiKeysNameRotateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(403)

	return json.NewEncoder(w).Encode(response)
}

type PostApiKeysNameRotate404Response struct {
}

func (response PostApiKeysNameRotate404Response) VisitPostApiKeysNameRotateResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetFlagsRequestObject struct {
}

type GetFlagsResponseObject interface {
	VisitGetFlagsResponse(w http.ResponseWriter) error
}

type GetFlags200JSONResponse []Flag

func (response GetFlags200JSONResponse) VisitGetFlagsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostFlagsRequestObject struct {
	Body *PostFlagsJSONRequestBody
}

type PostFlagsResponseObject interface {
	VisitPostFlagsResponse(w http.ResponseWriter) error
}

type PostFlags201JSONResponse Flag

func (response PostFlags201JSONResponse) VisitPostFlagsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type PostFlags400JSONResponse map[string]interface{}

func (response PostFlags400JSONResponse) VisitPostFlagsResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PostFlags409Response struct {
}

func (response PostFlags409Response) VisitPostFlagsResponse(w http.ResponseWriter) error {
	w.WriteHeader(409)
	return nil
}

type DeleteFlagsKeyRequestObject struct {
	Key string `json:"key"`
}

type DeleteFlagsKeyResponseObject interface {
	VisitDeleteFlagsKeyResponse(w http.ResponseWriter) error
}

type DeleteFlagsKey204Response struct {
}

func (response DeleteFlagsKey204Response) VisitDeleteFlagsKeyResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteFlagsKey404Response struct {
}

func (response DeleteFlagsKey404Response) VisitDeleteFlagsKeyResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetFlagsKeyRequestObject struct {
	Key string `json:"key"`
}

type GetFlagsKeyResponseObject interface {
	VisitGetFlagsKeyResponse(w http.ResponseWriter) error
}

type GetFlagsKey200JSONResponse Flag

func (response GetFlagsKey200JSONResponse) VisitGetFlagsKeyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetFlagsKey403Response struct {
}

func (response GetFlagsKey403Response) VisitGetFlagsKeyResponse(w http.ResponseWriter) error {
	w.WriteHeader(403)
	return nil
}

type GetFlagsKey404Response struct {
}

func (response GetFlagsKey404Response) VisitGetFlagsKeyResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PutFlagsKeyRequestObject struct {
	Key  string `json:"key"`
	Body *PutFlagsKeyJSONRequestBody
}

type PutFlagsKeyResponseObject interface {
	VisitPutFlagsKeyResponse(w http.ResponseWriter) error
}

type PutFlagsKey200JSONResponse Flag

func (response PutFlagsKey200JSONResponse) VisitPutFlagsKeyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PutFlagsKey400JSONResponse map[string]interface{}

func (response PutFlagsKey400JSONResponse) VisitPutFlagsKeyResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)

	return json.NewEncoder(w).Encode(response)
}

type PutFlagsKey404Response struct {
}

func (response PutFlagsKey404Response) VisitPutFlagsKeyResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetHealthzRequestObject struct {
}

type GetHealthzResponseObject interface {
	VisitGetHealthzResponse(w http.ResponseWriter) error
}

type GetHealthz200JSONResponse HealthResponse

func (response GetHealthz200JSONResponse) VisitGetHealthzResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetHealthz503JSONResponse HealthResponse

func (response GetHealthz503JSONResponse) VisitGetHealthzResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(503)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Create a new API key
	// (POST /api-keys)
	PostApiKeys(ctx context.Context, request PostApiKeysRequestObject) (PostApiKeysResponseObject, error)
	// Rotate an API key
	// (POST /api-keys/{name}/rotate)
	PostApiKeysNameRotate(ctx context.Context, request PostApiKeysNameRotateRequestObject) (PostApiKeysNameRotateResponseObject, error)
	// Retrieve all feature flags
	// (GET /flags)
	GetFlags(ctx context.Context, request GetFlagsRequestObject) (GetFlagsResponseObject, error)
	// Create a new feature flag
	// (POST /flags)
	PostFlags(ctx context.Context, request PostFlagsRequestObject) (PostFlagsResponseObject, error)
	// Delete a feature flag
	// (DELETE /flags/{key})
	DeleteFlagsKey(ctx context.Context, request DeleteFlagsKeyRequestObject) (DeleteFlagsKeyResponseObject, error)
	// Retrieve a single feature flag by key
	// (GET /flags/{key})
	GetFlagsKey(ctx context.Context, request GetFlagsKeyRequestObject) (GetFlagsKeyResponseObject, error)
	// Update an existing feature flag
	// (PUT /flags/{key})
	PutFlagsKey(ctx context.Context, request PutFlagsKeyRequestObject) (PutFlagsKeyResponseObject, error)
	// API healthcheck
	// (GET /healthz)
	GetHealthz(ctx context.Context, request GetHealthzRequestObject) (GetHealthzResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// PostApiKeys operation middleware
func (sh *strictHandler) PostApiKeys(w http.ResponseWriter, r *http.Request) {
	var request PostApiKeysRequestObject

	var body PostApiKeysJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostApiKeys(ctx, request.(PostApiKeysRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostApiKeys")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostApiKeysResponseObject); ok {
		if err := validResponse.VisitPostApiKeysResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostApiKeysNameRotate operation middleware
func (sh *strictHandler) PostApiKeysNameRotate(w http.ResponseWriter, r *http.Request, name string) {
	var request PostApiKeysNameRotateRequestObject

	request.Name = name

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostApiKeysNameRotate(ctx, request.(PostApiKeysNameRotateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostApiKeysNameRotate")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostApiKeysNameRotateResponseObject); ok {
		if err := validResponse.VisitPostApiKeysNameRotateResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetFlags operation middleware
func (sh *strictHandler) GetFlags(w http.ResponseWriter, r *http.Request) {
	var request GetFlagsRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetFlags(ctx, request.(GetFlagsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetFlags")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetFlagsResponseObject); ok {
		if err := validResponse.VisitGetFlagsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostFlags operation middleware
func (sh *strictHandler) PostFlags(w http.ResponseWriter, r *http.Request) {
	var request PostFlagsRequestObject

	var body PostFlagsJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostFlags(ctx, request.(PostFlagsRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostFlags")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostFlagsResponseObject); ok {
		if err := validResponse.VisitPostFlagsResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteFlagsKey operation middleware
func (sh *strictHandler) DeleteFlagsKey(w http.ResponseWriter, r *http.Request, key string) {
	var request DeleteFlagsKeyRequestObject

	request.Key = key

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteFlagsKey(ctx, request.(DeleteFlagsKeyRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteFlagsKey")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteFlagsKeyResponseObject); ok {
		if err := validResponse.VisitDeleteFlagsKeyResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetFlagsKey operation middleware
func (sh *strictHandler) GetFlagsKey(w http.ResponseWriter, r *http.Request, key string) {
	var request GetFlagsKeyRequestObject

	request.Key = key

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetFlagsKey(ctx, request.(GetFlagsKeyRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetFlagsKey")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetFlagsKeyResponseObject); ok {
		if err := validResponse.VisitGetFlagsKeyResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PutFlagsKey operation middleware
func (sh *strictHandler) PutFlagsKey(w http.ResponseWriter, r *http.Request, key string) {
	var request PutFlagsKeyRequestObject

	request.Key = key

	var body PutFlagsKeyJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PutFlagsKey(ctx, request.(PutFlagsKeyRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutFlagsKey")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PutFlagsKeyResponseObject); ok {
		if err := validResponse.VisitPutFlagsKeyResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetHealthz operation middleware
func (sh *strictHandler) GetHealthz(w http.ResponseWriter, r *http.Request) {
	var request GetHealthzRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetHealthz(ctx, request.(GetHealthzRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetHealthz")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetHealthzResponseObject); ok {
		if err := validResponse.VisitGetHealthzResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
