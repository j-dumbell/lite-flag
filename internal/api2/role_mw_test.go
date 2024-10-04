package api2

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRoleMW(t *testing.T) {
	resetDB(t)
	roleMW := newRoleMW(authService)
	key := createAdminKey(t)

	var requestCtx context.Context
	handler := http.HandlerFunc(func(_ http.ResponseWriter, r *http.Request) {
		requestCtx = r.Context()
	})

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	require.NoError(t, err, "request failed")
	req.Header.Add(ApiKeyHeader, key.Key)

	rr := httptest.NewRecorder()
	roleMW(handler).ServeHTTP(rr, req)

	user, ok := getUser(requestCtx)
	require.True(t, ok, "ok")
	assert.Equal(t, key.Redact(), user, "user")
}
