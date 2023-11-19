package api

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/internal/health"
	"github.com/j-dumbell/lite-flag/internal/key"
)

func New(db *sql.DB) *chi.Mux {
	flagRepo := fflag.NewRepo(db)
	keyRepo := key.NewRepo(db)

	healthHander := health.NewHandler(db)
	flagHandler := fflag.NewHandler(&flagRepo)
	keyHandler := key.NewHandler(&keyRepo)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", healthHander.Get)

	r.Get("/flags", flagHandler.Get)
	r.Post("/flags", flagHandler.Post)
	r.Get("/flags/{flagID}", flagHandler.GetOne)
	r.Delete("/flags/{flagID}", flagHandler.Delete)

	r.Get("/api-keys", keyHandler.Get)
	r.Post("/api-keys", keyHandler.Post)
	r.Delete("/api-keys/{keyID}", keyHandler.Delete)
	return r
}
