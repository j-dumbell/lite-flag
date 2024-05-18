package api

import (
	"net/http"

	"github.com/j-dumbell/lite-flag/pkg/chix"
)

type healthResponse struct {
	Database bool `json:"Database"`
}

func (api *API) Healthcheck(r *http.Request) chix.Response {
	err := api.db.Ping()
	if err != nil {
		response := healthResponse{Database: false}
		return chix.NotFound(response)
	}
	return chix.NoContent()
}
