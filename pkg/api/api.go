package api

import (
	"database/sql"
	"github.com/j-dumbell/lite-flag/pkg/auth"
	"github.com/j-dumbell/lite-flag/pkg/health"
	"github.com/j-dumbell/lite-flag/pkg/key"
	"net/http"
)

func New(db *sql.DB) *http.ServeMux {
	authMiddleware := auth.NewMiddleware(key.NewRepo(db))
	keyHandler := key.NewHandler(db)

	mux := http.NewServeMux()
	mux.Handle("/health", health.NewHandler(db))
	mux.Handle("/api-keys", authMiddleware.Wrap(keyHandler, []key.Role{key.Root, key.Admin}))
	mux.Handle("/api-keys/", keyHandler)

	return mux
}
