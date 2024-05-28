package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusNoContent, res.StatusCode)
}
