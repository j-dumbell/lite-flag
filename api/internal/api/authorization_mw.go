package api

import (
	"net/http"
	"slices"

	"github.com/j-dumbell/lite-flag/internal/auth"
)

func authMW(permittedRoles ...auth.Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			user, ok := getUser(r.Context())
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if !slices.Contains(permittedRoles, user.Role) {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

var adminOnly = authMW(auth.RoleRoot, auth.RoleAdmin)
