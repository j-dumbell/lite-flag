package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/j-dumbell/lite-flag/internal/auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostKey(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	reqBody := auth.CreateApiKeyParams{
		Name: "some-key",
		Role: auth.RoleReadonly,
	}
	jsonReqBody, err := json.Marshal(reqBody)
	require.NoError(t, err, "could not marshal request reqBody")

	req := httptest.NewRequest(http.MethodPost, "/api-keys", bytes.NewReader(jsonReqBody))
	req.Header.Add(ApiKeyHeader, key.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	resultBody := result.Body
	defer resultBody.Close()
	var actualBody auth.ApiKey
	err = json.NewDecoder(resultBody).Decode(&actualBody)
	require.NoError(t, err, "could not decode response reqBody")

	assert.Equal(t, http.StatusCreated, result.StatusCode, "status code")
	assert.Equal(t, actualBody.Name, reqBody.Name, "Key")
	assert.Equal(t, actualBody.Role, reqBody.Role, "Role")
	assert.GreaterOrEqual(t, len(actualBody.Key), 40, "Key length")
}

func TestDeleteKey(t *testing.T) {
	resetDB(t)
	requestorKey := createAdminKey(t)

	apiKey, err := authService.CreateKey(context.Background(), auth.CreateApiKeyParams{
		Name: "blah",
		Role: auth.RoleReadonly,
	})
	require.NoError(t, err, "failed to insert test key")

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api-keys/%s", apiKey.Name), nil)
	req.Header.Add(ApiKeyHeader, requestorKey.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()

	assert.Equal(t, http.StatusOK, result.StatusCode, "status code")
}

func TestDeleteKey_adminDeletingAdmin(t *testing.T) {
	resetDB(t)
	requestorKey := createAdminKey(t)

	apiKey, err := authService.CreateKey(context.Background(), auth.CreateApiKeyParams{
		Name: "blah",
		Role: auth.RoleAdmin,
	})
	require.NoError(t, err, "failed to insert test key")

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api-keys/%s", apiKey.Name), nil)
	req.Header.Add(ApiKeyHeader, requestorKey.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()

	assert.Equal(t, http.StatusForbidden, result.StatusCode, "status code")
}

func TestDeleteKey_root(t *testing.T) {
	resetDB(t)
	requestorKey := createAdminKey(t)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api-keys/%s", requestorKey.Name), nil)
	req.Header.Add(ApiKeyHeader, requestorKey.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()

	assert.Equal(t, http.StatusForbidden, result.StatusCode, "status code")
}

func TestRotateKey(t *testing.T) {
	resetDB(t)
	requestorKey := createAdminKey(t)

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api-keys/%s/rotate", requestorKey.Name), nil)
	req.Header.Add(ApiKeyHeader, requestorKey.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	resultBody := result.Body
	defer resultBody.Close()
	var actualBody auth.ApiKey
	err := json.NewDecoder(resultBody).Decode(&actualBody)
	require.NoError(t, err, "could not decode response response body")

	assert.Equal(t, http.StatusOK, result.StatusCode, "status code")
	assert.NotEqual(t, requestorKey.Key, actualBody.Key, "rotated key should not equal original key")
	assert.Equal(t, requestorKey.Name, actualBody.Name, "Key")
	assert.Equal(t, requestorKey.Role, actualBody.Role, "Role")
}

func TestRotateKey_notFound(t *testing.T) {
	resetDB(t)
	requestorKey := createRootKey(t)

	req := httptest.NewRequest(http.MethodPost, "/api-keys/1000/rotate", nil)
	req.Header.Add(ApiKeyHeader, requestorKey.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusNotFound, result.StatusCode, "status code")
}

func TestRotateKey_adminRotatingAnothersKey(t *testing.T) {
	resetDB(t)
	requestorKey := createAdminKey(t)

	apiKey, err := authService.CreateKey(context.Background(), auth.CreateApiKeyParams{
		Name: "blah",
		Role: auth.RoleAdmin,
	})
	require.NoError(t, err, "failed to insert test key")

	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/api-keys/%s/rotate", apiKey.Name), nil)
	req.Header.Add(ApiKeyHeader, requestorKey.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusForbidden, result.StatusCode, "status code")
}
