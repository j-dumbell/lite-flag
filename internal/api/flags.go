package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/chix"
	"github.com/j-dumbell/lite-flag/pkg/validation"
)

func (api *API) PostFlag(r *http.Request) chix.Response {
	var body fflag.Flag
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return chix.BadRequest("invalid JSON body")
	}

	flag, err := api.flagService.Create(r.Context(), body)
	if errors.As(err, &validation.Result{}) {
		return chix.BadRequest(err)
	} else if errors.Is(err, fflag.ErrAlreadyExists) {
		return chix.Conflict("a flag with that name already exists")
	} else if err != nil {
		return chix.InternalServerError()
	}

	return chix.Created(flag)
}

type PutFlagBody struct {
	Type         fflag.FlagType         `json:"type"`
	BooleanValue *bool                  `json:"booleanValue"`
	StringValue  *string                `json:"stringValue"`
	JSONValue    map[string]interface{} `json:"jsonValue"`
}

func (api *API) PutFlag(r *http.Request) chix.Response {
	key := chi.URLParam(r, "key")

	var body PutFlagBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return chix.BadRequest("invalid JSON body")
	}

	flag := fflag.Flag{
		Key:          key,
		Type:         body.Type,
		BooleanValue: body.BooleanValue,
		StringValue:  body.StringValue,
		JSONValue:    body.JSONValue,
	}

	flag, err := api.flagService.Update(r.Context(), flag)
	if errors.As(err, &validation.Result{}) {
		return chix.BadRequest(err)
	} else if errors.Is(err, fflag.ErrNotFound) {
		return chix.NotFound(nil)
	} else if err != nil {
		return chix.InternalServerError()
	}

	return chix.Created(flag)
}

func (api *API) GetFlags(r *http.Request) chix.Response {
	flags, err := api.flagService.FindAll(r.Context())
	if err != nil {
		return chix.InternalServerError()
	}

	return chix.OK(flags)
}

func (api *API) GetFlag(r *http.Request) chix.Response {
	key := chi.URLParam(r, "key")

	flag, err := api.flagService.FindOne(r.Context(), key)
	if errors.Is(err, fflag.ErrNotFound) {
		return chix.NotFound("a flag with that ID does not exist")
	} else if err != nil {
		return chix.InternalServerError()
	}

	return chix.OK(flag)
}

func (api *API) DeleteFlag(r *http.Request) chix.Response {
	key := chi.URLParam(r, "key")

	err := api.flagService.Delete(r.Context(), key)
	if err == fflag.ErrNotFound {
		return chix.NotFound(nil)
	} else if err != nil {
		return chix.InternalServerError()
	}

	return chix.OK(nil)
}
