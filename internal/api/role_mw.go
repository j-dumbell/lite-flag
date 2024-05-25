package api

import (
	"context"
	"net/http"

	"github.com/j-dumbell/lite-flag/internal/auth"
)

type ctxKey int

const ctxKeyRole ctxKey = 0
const apiKeyHeader = "X-Api-Key"

func getRole(ctx context.Context) auth.Role {
	roleValue := ctx.Value(ctxKeyRole)
	role, ok := roleValue.(auth.Role)
	if !ok {
		return ""
	}
	return role
}

// newRoleMW is a middleware which reads the API key provided in the headers
// and writes the corresponding role to the request context if valid.
func newRoleMW(authService auth.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get(apiKeyHeader)
			if apiKey == "" {
				next.ServeHTTP(w, r)
				return
			}

			role, err := authService.KeyRole(apiKey)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), ctxKeyRole, role)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
