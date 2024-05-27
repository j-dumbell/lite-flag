package api

import (
	"database/sql"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/j-dumbell/lite-flag/internal/auth"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/chix"
)

var requestTimeout = 20 * time.Second

type API struct {
	db          *sql.DB
	flagService fflag.Service
	authService auth.Service
}

func New(db *sql.DB, flagService fflag.Service, authService auth.Service) API {
	return API{
		db:          db,
		flagService: flagService,
		authService: authService,
	}
}

func (api *API) NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(requestTimeout))
	r.Use(newRoleMW(api.authService))

	chix.Get(r, "/healthz", api.Healthcheck)

	chix.Get(r, "/flags", api.GetFlags, anyRole)
	chix.Post(r, "/flags", api.PostFlag, adminOnly)
	chix.Get(r, "/flags/{id}", api.GetFlag)
	chix.Delete(r, "/flags/{id}", api.DeleteFlag, adminOnly)

	chix.Post(r, "/api-keys", api.PostKey, adminOnly)
	chix.Delete(r, "/api-keys/{id}", api.DeleteKey, adminOnly)

	return r
}
