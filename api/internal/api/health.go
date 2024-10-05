package api

import (
	"context"

	"github.com/j-dumbell/lite-flag/internal/oapi"
)

func (srv *server) GetHealthz(ctx context.Context, _ oapi.GetHealthzRequestObject) (oapi.GetHealthzResponseObject, error) {
	err := srv.db.PingContext(ctx)
	if err != nil {
		response := oapi.GetHealthz503JSONResponse(oapi.HealthResponse{
			Database: false,
		})
		return response, nil
	}

	response := oapi.GetHealthz200JSONResponse(oapi.HealthResponse{
		Database: true,
	})

	return response, nil
}
