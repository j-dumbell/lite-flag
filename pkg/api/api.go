package api

import (
	"database/sql"
	"github.com/j-dumbell/lite-flag/pkg/auth"
	"github.com/j-dumbell/lite-flag/pkg/fflag"
	"github.com/j-dumbell/lite-flag/pkg/health"
	"github.com/j-dumbell/lite-flag/pkg/key"
	"net/http"
)

func New(db *sql.DB) *http.ServeMux {
	authMiddleware := auth.NewMiddleware(key.NewRepo(db))

	flagRepo := fflag.NewRepo(db)
	keyRepo := key.NewRepo(db)

	keyHandler := key.NewHandler(keyRepo)
	flagHandler := fflag.NewHandler(flagRepo)

	mux := http.NewServeMux()
	mux.Handle("/health", health.NewHandler(db))
	mux.Handle("/api-keys", authMiddleware.Wrap(keyHandler, []key.Role{key.Root, key.Admin}))
	mux.Handle("/api-keys/", keyHandler)
	mux.Handle("/flags", authMiddleware.Wrap(flagHandler, []key.Role{key.Root, key.Admin}))

	return mux
}
