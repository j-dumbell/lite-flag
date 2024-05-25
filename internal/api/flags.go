package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/chix"
	"github.com/j-dumbell/lite-flag/pkg/pg"
	"github.com/j-dumbell/lite-flag/pkg/validation"
)

func (api *API) PostFlag(r *http.Request) chix.Response {
	var body fflag.UpsertFlagParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return chix.BadRequest("invalid JSON body")
	}

	flag, err := api.flagService.Create(body)
	if errors.Is(err, pg.ErrAlreadyExists) {
		return chix.Conflict("a flag with that name already exists")
	} else if errors.As(err, &validation.ValidationError{}) {
		return chix.BadRequest(err)
	}

	return chix.Created(flag)
}

func (api *API) GetFlags(r *http.Request) chix.Response {
	flags, err := api.flagService.FindAll()
	if err != nil {
		return chix.InternalServerError()
	}

	return chix.OK(flags)
}

func (api *API) GetFlag(r *http.Request) chix.Response {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return chix.NotFound("a flag with that ID does not exist")
	}

	flag, err := api.flagService.FindOne(uint32(id))
	if errors.Is(err, pg.ErrNoRows) {
		return chix.NotFound("a flag with that ID does not exist")
	} else if err != nil {
		return chix.InternalServerError()
	}

	return chix.OK(flag)
}

func (api *API) DeleteFlag(r *http.Request) chix.Response {
	idParam := chi.URLParam(r, "id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return chix.NotFound("a flag with that ID does not exist")
	}

	err = api.flagService.Delete(uint32(id))
	if err != nil {
		return chix.InternalServerError()
	}

	return chix.OK(nil)
}
