package auth

import (
	"database/sql"
	"github.com/j-dumbell/lite-flag/pkg/array"
	"github.com/j-dumbell/lite-flag/pkg/key"
	"net/http"
)

type ApiKeyRepo interface {
	FindOneByKey(key string) (key.ApiKey, error)
}

type Middleware struct {
	apiKeyRepo ApiKeyRepo
}

func NewMiddleware(apiKeyRepo ApiKeyRepo) Middleware {
	return Middleware{apiKeyRepo: apiKeyRepo}
}

func (middleware Middleware) Wrap(handler http.Handler, roles []key.Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerKey := r.Header.Get("x-api-key")
		if headerKey == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		apiKey, err := middleware.apiKeyRepo.FindOneByKey(headerKey)
		if err != nil {
			switch err {
			case sql.ErrNoRows:
				w.WriteHeader(http.StatusUnauthorized)
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		if !array.Includes(roles, apiKey.Role) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		handler.ServeHTTP(w, r)
	})
}
