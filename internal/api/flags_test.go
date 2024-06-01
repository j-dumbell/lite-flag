package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/pkg/fp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFlags(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

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
	req.Header.Add(ApiKeyHeader, key.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	resultBody := result.Body
	defer resultBody.Close()
	var actual []fflag.Flag
	err = json.NewDecoder(resultBody).Decode(&actual)
	require.NoError(t, err, "could not decode response body")

	expected := []fflag.Flag{flag1, flag2}

	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, 2, len(actual), "unexpected number of flags in response")
	assert.Subset(t, expected, actual, "unexpected flags in response")
}

func TestGetFlag(t *testing.T) {
	resetDB(t)

	savedFlag, err := flagService.Create(context.Background(), fflag.Flag{
		Key:          "blah",
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     true,
		BooleanValue: fp.ToPtr(false),
	})
	require.NoError(t, err, "could not save test data to DB")

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/flags/%s", savedFlag.Key), nil)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	resultBody := result.Body
	defer resultBody.Close()
	var actual fflag.Flag
	err = json.NewDecoder(resultBody).Decode(&actual)
	require.NoError(t, err, "could not decode response body")

	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, savedFlag, actual)
}

func TestGetFlag_notPublic_authorized(t *testing.T) {
	resetDB(t)
	key := createReadonlyKey(t)

	savedFlag, err := flagService.Create(context.Background(), fflag.Flag{
		Key:         "blah",
		Type:        fflag.FlagTypeString,
		IsPublic:    false,
		StringValue: fp.ToPtr("yoyo"),
	})
	require.NoError(t, err, "could not save test data to DB")

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/flags/%s", savedFlag.Key), nil)
	req.Header.Add(ApiKeyHeader, key.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	resultBody := result.Body
	defer resultBody.Close()
	var actual fflag.Flag
	err = json.NewDecoder(resultBody).Decode(&actual)
	require.NoError(t, err, "could not decode response body")

	assert.Equal(t, savedFlag, actual)
}

func TestGetFlag_notPublic_unauthorized(t *testing.T) {
	resetDB(t)

	savedFlag, err := flagService.Create(context.Background(), fflag.Flag{
		Key:         "blah",
		Type:        fflag.FlagTypeString,
		IsPublic:    false,
		StringValue: fp.ToPtr("yoyo"),
	})
	require.NoError(t, err, "could not save test data to DB")

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/flags/%s", savedFlag.Key), nil)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusUnauthorized, result.StatusCode)
}

func TestPostFlag(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	reqBody := fflag.Flag{
		Key:          "my-flag",
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     true,
		BooleanValue: fp.ToPtr(false),
	}
	jsonReqBody, err := json.Marshal(reqBody)
	require.NoError(t, err, "could not marshal request body")

	req := httptest.NewRequest(http.MethodPost, "/flags", bytes.NewReader(jsonReqBody))
	req.Header.Add(ApiKeyHeader, key.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	resultBody := result.Body
	defer resultBody.Close()
	var actualBody fflag.Flag
	err = json.NewDecoder(resultBody).Decode(&actualBody)
	require.NoError(t, err, "could not decode response body")

	assert.Equal(t, http.StatusCreated, result.StatusCode, "status code")
	assert.Equal(t, actualBody.Key, reqBody.Key, "Key")
	assert.Equal(t, actualBody.Type, reqBody.Type, "Type")
	assert.Equal(t, actualBody.BooleanValue, reqBody.BooleanValue, "BooleanValue")
	assert.Equal(t, actualBody.StringValue, reqBody.StringValue, "StringValue")
	assert.Equal(t, actualBody.JSONValue, reqBody.JSONValue, "JSONValue")
}

func TestPostFlag_alreadyExists(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	flagName := "some-flag"
	_, err := flagService.Create(context.Background(), fflag.Flag{
		Key:          flagName,
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     true,
		BooleanValue: fp.ToPtr(false),
	})

	flag := fflag.Flag{
		Key:         flagName,
		Type:        fflag.FlagTypeString,
		IsPublic:    true,
		StringValue: fp.ToPtr("abc"),
	}
	jsonReqBody, err := json.Marshal(flag)
	require.NoError(t, err, "could not marshal request body")

	req := httptest.NewRequest(http.MethodPost, "/flags", bytes.NewReader(jsonReqBody))
	req.Header.Add(ApiKeyHeader, key.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusConflict, result.StatusCode, "status code")
}

func TestPostFlag_invalidBody(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	req := httptest.NewRequest(http.MethodPost, "/flags", bytes.NewReader([]byte(`{enabled: false}`)))
	req.Header.Add(ApiKeyHeader, key.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusBadRequest, result.StatusCode, "status code")
}

func TestDeleteFlag(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	flag, err := flagService.Create(context.Background(), fflag.Flag{
		Key:          "fooBar",
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     true,
		BooleanValue: fp.ToPtr(true),
	})
	require.NoError(t, err, "could not save test data to DB")

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/flags/%s", flag.Key), nil)
	req.Header.Add(ApiKeyHeader, key.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)
}

func TestDeleteFlag_notFound(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	req := httptest.NewRequest(http.MethodDelete, "/flags/100", nil)
	req.Header.Add(ApiKeyHeader, key.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusNotFound, result.StatusCode)
}

func TestPutFlag(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	savedFlag1 := fflag.Flag{
		Key:         "abc",
		Type:        fflag.FlagTypeString,
		IsPublic:    true,
		StringValue: fp.ToPtr("foo-bar"),
	}
	_, err := flagService.Create(context.Background(), savedFlag1)
	require.NoError(t, err, "could not setup test data")

	reqBody, err := json.Marshal(PutFlagBody{
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     false,
		BooleanValue: fp.ToPtr(false),
	})
	require.NoError(t, err, "failed to marshal req body")

	req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/flags/%s", savedFlag1.Key), bytes.NewReader(reqBody))
	req.Header.Add(ApiKeyHeader, key.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	resultBody := result.Body
	defer resultBody.Close()
	var actual fflag.Flag

	err = json.NewDecoder(resultBody).Decode(&actual)
	require.NoError(t, err, "could not decode response body")

	expected := fflag.Flag{
		Key:          savedFlag1.Key,
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     false,
		BooleanValue: fp.ToPtr(false),
	}
	assert.Equal(t, expected, actual)
}

func TestPutFlag_notFound(t *testing.T) {
	resetDB(t)
	key := createAdminKey(t)

	reqBody, err := json.Marshal(PutFlagBody{
		Type:         fflag.FlagTypeBoolean,
		IsPublic:     true,
		BooleanValue: fp.ToPtr(false),
	})
	require.NoError(t, err, "failed to marshal req body")

	req := httptest.NewRequest(http.MethodPut, "/flags/invalid-key", bytes.NewReader(reqBody))
	req.Header.Add(ApiKeyHeader, key.Key)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusNotFound, result.StatusCode)
}
