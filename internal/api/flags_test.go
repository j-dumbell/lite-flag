package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/j-dumbell/lite-flag/internal/bootstrap"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFlags(t *testing.T) {
	err := bootstrap.Truncate(testDB)
	require.NoError(t, err, "failed to refresh DB")

	savedFlag1 := fflag.UpsertFlagParams{
		Name:    "abc",
		Enabled: true,
	}
	flag1, err := flagService.Create(savedFlag1)
	require.NoError(t, err, "could not setup test data")

	savedFlag2 := fflag.UpsertFlagParams{
		Name:    "def",
		Enabled: true,
	}
	flag2, err := flagService.Create(savedFlag2)
	require.NoError(t, err, "could not setup test data")

	req := httptest.NewRequest(http.MethodGet, "/flags", nil)
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
	err := bootstrap.Truncate(testDB)
	require.NoError(t, err, "failed to refresh DB")

	savedFlag, err := flagService.Create(fflag.UpsertFlagParams{
		Name:    "blah",
		Enabled: false,
	})
	require.NoError(t, err, "could not save test data to DB")

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/flags/%d", savedFlag.ID), nil)
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

func TestPostFlag(t *testing.T) {
	err := bootstrap.Truncate(testDB)
	require.NoError(t, err, "failed to refresh DB")

	reqBody := fflag.UpsertFlagParams{
		Name:    "my-flag",
		Enabled: false,
	}
	jsonReqBody, err := json.Marshal(reqBody)
	require.NoError(t, err, "could not marshal request body")

	req := httptest.NewRequest(http.MethodPost, "/flags", bytes.NewReader(jsonReqBody))
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	resultBody := result.Body
	defer resultBody.Close()
	var actualBody fflag.Flag
	err = json.NewDecoder(resultBody).Decode(&actualBody)
	require.NoError(t, err, "could not decode response body")

	assert.Equal(t, http.StatusCreated, result.StatusCode, "status code")
	assert.Equal(t, actualBody.Name, reqBody.Name, "Name")
	assert.Equal(t, actualBody.Enabled, reqBody.Enabled, "Enabled")
}

func TestDeleteFlag(t *testing.T) {
	err := bootstrap.Truncate(testDB)
	require.NoError(t, err, "failed to refresh DB")

	flag, err := flagService.Create(fflag.UpsertFlagParams{
		Name:    "fooBar",
		Enabled: true,
	})
	require.NoError(t, err, "could not save test data to DB")

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/flags/%d", flag.ID), nil)
	w := httptest.NewRecorder()
	testApi.NewRouter().ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)
}
