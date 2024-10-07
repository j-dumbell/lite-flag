package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()
	testHandler.ServeHTTP(w, req)

	res := w.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode, "status code")

	var actualResBody map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&actualResBody)
	require.NoError(t, err, "failed to decode response body")

	expectedResBody := map[string]interface{}{"database": true}
	assert.Equal(t, expectedResBody, actualResBody)
}
