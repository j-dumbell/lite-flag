package api2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/internal/oapi"
	"github.com/j-dumbell/lite-flag/pkg/fp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFlags(t *testing.T) {
	resetDB(t)

	savedFlag1 := fflag.Flag{
		Key:          "abc",
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     true,
		BooleanValue: fp.ToPtr(true),
	}
	flag1, err := flagService.Create(context.Background(), savedFlag1)
	require.NoError(t, err, "could not setup test data")

	savedFlag2 := fflag.Flag{
		Key:          "def",
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     false,
		BooleanValue: fp.ToPtr(true),
	}
	flag2, err := flagService.Create(context.Background(), savedFlag2)
	require.NoError(t, err, "could not setup test data")

	req := httptest.NewRequest(http.MethodGet, "/flags", nil)
	w := httptest.NewRecorder()
	testHandler.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	resultBody := result.Body
	defer resultBody.Close()
	var actual []oapi.Flag
	err = json.NewDecoder(resultBody).Decode(&actual)
	require.NoError(t, err, "could not decode response body")

	expected := fp.Map([]fflag.Flag{flag1, flag2}, toFlagDTO)

	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.ElementsMatch(t, expected, actual, "unexpected flags in response")
}

func TestPostFlag(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	reqBody := oapi.Flag{
		Key:      "my-flag",
		Type:     oapi.FlagTypeBoolean,
		IsPublic: true,
	}
	_ = reqBody.Value.FromFlagValue0(true)

	jsonReqBody, err := json.Marshal(reqBody)
	require.NoError(t, err, "could not marshal request body")

	req := httptest.NewRequest(http.MethodPost, "/flags", bytes.NewReader(jsonReqBody))
	req.Header.Add(ApiKeyHeader, key.Key)

	w := httptest.NewRecorder()
	testHandler.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusCreated, result.StatusCode, "status code")

	resultBody := result.Body
	defer resultBody.Close()
	var actualBody oapi.Flag
	err = json.NewDecoder(resultBody).Decode(&actualBody)
	require.NoError(t, err, "could not decode response body")
	assert.Equal(t, actualBody, reqBody, "response body")
}

func TestPostFlag_badRequest(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	req := httptest.NewRequest(http.MethodPost, "/flags", nil)
	req.Header.Add(ApiKeyHeader, key.Key)

	w := httptest.NewRecorder()
	testHandler.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusBadRequest, result.StatusCode, "status code")
}

func TestDeleteFlag(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	flag, err := flagService.Create(context.Background(), fflag.Flag{
		Key:         "abcd-e",
		Type:        fflag.FlagTypeString,
		IsPublic:    false,
		StringValue: fp.ToPtr("flag value"),
	})
	require.NoError(t, err, "could not setup test data")

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/flags/%s", flag.Key), nil)
	req.Header.Add(ApiKeyHeader, key.Key)

	w := httptest.NewRecorder()
	testHandler.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusNoContent, result.StatusCode, "status code")
}

func TestDeleteFlag_notFound(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	req := httptest.NewRequest(http.MethodDelete, "/flags/abc", nil)
	req.Header.Add(ApiKeyHeader, key.Key)

	w := httptest.NewRecorder()
	testHandler.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusNotFound, result.StatusCode, "status code")
}

func TestGetFlag(t *testing.T) {
	resetDB(t)

	flag := fflag.Flag{
		Key:          "abc-123",
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     true,
		BooleanValue: fp.ToPtr(true),
	}
	savedFlag, err := flagService.Create(context.Background(), flag)
	require.NoError(t, err, "could not setup test data")

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/flags/%s", flag.Key), nil)

	w := httptest.NewRecorder()
	testHandler.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode, "status code")

	resultBody := result.Body
	defer resultBody.Close()
	var actual oapi.Flag
	err = json.NewDecoder(resultBody).Decode(&actual)
	require.NoError(t, err, "could not decode response body")
	assert.Equal(t, toFlagDTO(savedFlag), actual, "response body")
}

func TestGetFlag_notFound(t *testing.T) {
	resetDB(t)

	req := httptest.NewRequest(http.MethodGet, "/flags/uhoh", nil)

	w := httptest.NewRecorder()
	testHandler.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusNotFound, result.StatusCode, "status code")
}

func TestPutFlag(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	flag := fflag.Flag{
		Key:          "abc-123",
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     false,
		BooleanValue: fp.ToPtr(false),
	}
	savedFlag, err := flagService.Create(context.Background(), flag)
	require.NoError(t, err, "could not setup test data")

	updateFlagBody := oapi.PutFlagsKeyJSONRequestBody{
		IsPublic: true,
		Type:     oapi.FlagInputTypeString,
	}
	updateFlagBody.Value.FromFlagInputValue1("xyz")

	updateFlagBodyBytes, err := json.Marshal(updateFlagBody)
	require.NoError(t, err, "failed to marshal request body")

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/flags/%s", savedFlag.Key), bytes.NewReader(updateFlagBodyBytes))
	req.Header.Add(ApiKeyHeader, key.Key)

	w := httptest.NewRecorder()
	testHandler.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode, "status code")

	// ToDo - assert on body
}

func TestPutFlag_notFound(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	updateFlagBody := oapi.PutFlagsKeyJSONRequestBody{
		IsPublic: true,
		Type:     oapi.FlagInputTypeString,
	}
	updateFlagBody.Value.FromFlagInputValue1("xyz")

	updateFlagBodyBytes, err := json.Marshal(updateFlagBody)
	require.NoError(t, err, "failed to marshal request body")

	req := httptest.NewRequest(http.MethodPut, "/flags/notfound", bytes.NewReader(updateFlagBodyBytes))
	req.Header.Add(ApiKeyHeader, key.Key)

	w := httptest.NewRecorder()
	testHandler.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusNotFound, result.StatusCode, "status code")
}
