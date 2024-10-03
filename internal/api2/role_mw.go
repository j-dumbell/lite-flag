package api2

import (
	"context"
	"net/http"

	"github.com/j-dumbell/lite-flag/internal/auth"
)

type ctxKey int

const ctxKeyUser ctxKey = 0

const ApiKeyHeader = "X-Api-Key" //nolint:gosec

func getUser(ctx context.Context) (auth.ApiKeyRedacted, bool) {
	roleValue := ctx.Value(ctxKeyUser)
	apiKeyRedacted, ok := roleValue.(auth.ApiKeyRedacted)
	if !ok {
		return auth.ApiKeyRedacted{}, false
	}
	return apiKeyRedacted, true
}

// newRoleMW is a middleware which reads the API key provided in the headers
// and writes the corresponding role to the request context if valid.
func newRoleMW(authService auth.Service) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get(ApiKeyHeader)
			if apiKey == "" {
				next.ServeHTTP(w, r)
				return
			}

			apiKeyRedacted, err := authService.FindOneByKey(r.Context(), apiKey)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), ctxKeyUser, apiKeyRedacted)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
