package api

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/chix"
)

type API struct {
	db          *sql.DB
	flagService fflag.Service
}

func New(db *sql.DB, flagService fflag.Service) API {
	return API{
		db:          db,
		flagService: flagService,
	}
}

func (api *API) NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	chix.RegisterGet(r, "/healthz", api.Healthcheck)

	chix.RegisterGet(r, "/flags", api.GetFlags)
	chix.RegisterPost(r, "/flags", api.PostFlag)
	chix.RegisterGet(r, "/flags/{id}", api.GetFlag)
	chix.RegisterDelete(r, "/flags/{id}", api.DeleteFlag)

	return r
}
