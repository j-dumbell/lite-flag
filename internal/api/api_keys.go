package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/j-dumbell/lite-flag/internal/auth"
	"github.com/j-dumbell/lite-flag/pkg/chix"
	"github.com/j-dumbell/lite-flag/pkg/validation"
)

func (api *API) PostKey(r *http.Request) chix.Response {
	var body auth.CreateApiKeyParams
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return chix.BadRequest("invalid JSON body")
	}

	user, _ := getUser(r.Context())
	if user.Role == auth.RoleAdmin && body.Role != auth.RoleReadonly {
		return chix.Forbidden("admins may only create API keys with readonly role")
	}

	apiKey, err := api.authService.CreateKey(body)
	if errors.As(err, &validation.Result{}) {
		return chix.BadRequest(err)
	} else if errors.Is(err, auth.ErrAlreadyExists) {
		return chix.Conflict(err)
	} else if err != nil {
		return chix.InternalServerError()
	}

	return chix.Created(apiKey)
}

func (api *API) DeleteKey(r *http.Request) chix.Response {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return chix.NotFound(nil)
	}

	apiKeyRedacted, err := api.authService.FindOneByID(id)
	if err == auth.ErrKeyNotFound {
		return chix.NotFound(nil)
	} else if err != nil {
		return chix.InternalServerError()
	}

	requestor, _ := getUser(r.Context())
	if apiKeyRedacted.Role == auth.RoleRoot {
		return chix.Forbidden("root API key cannot be deleted")
	}
	if requestor.Role == auth.RoleAdmin && apiKeyRedacted.Role != auth.RoleReadonly {
		return chix.Forbidden("admins may only delete API keys with readonly role")
	}

	if err := api.authService.DeleteByID(id); err != nil {
		return chix.InternalServerError()
	}

	return chix.OK(nil)
}

// func (api *API) RotateKey(r *http.Request) chix.Response {
// 	idParam := chi.URLParam(r, "id")
// 	id, err := strconv.Atoi(idParam)
// 	if err != nil {
// 		return chix.NotFound(nil)
// 	}
//
// 	apiKeyRedacted, err := api.authService.FindOneByID(id)
// 	if err == auth.ErrKeyNotFound {
// 		return chix.NotFound(nil)
// 	} else if err != nil {
// 		return chix.InternalServerError()
// 	}
//
// 	// ToDo - Need requestor ID in context
// 	requestor := getUser(r.Context())
// }
