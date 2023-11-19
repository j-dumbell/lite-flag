package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetFlags(t *testing.T) {
	flagRepo := fflag.NewRepo(testDB)

	savedFlag1 := fflag.Flag{
		ID:      "james-flag-1",
		Enabled: true,
	}
	err := flagRepo.Save(savedFlag1)
	require.NoError(t, err, "could not save test data to DB")

	savedFlag2 := fflag.Flag{
		ID:      "james-flag-2",
		Enabled: false,
	}
	err = flagRepo.Save(savedFlag2)
	require.NoError(t, err, "could not save test data to DB")

	req := httptest.NewRequest(http.MethodGet, "/flags", nil)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	resultBody := result.Body
	defer resultBody.Close()
	var actualBody fflag.GetResponse
	err = json.NewDecoder(resultBody).Decode(&actualBody)
	require.NoError(t, err, "could not decode response body")

	assert.Equal(t, 2, len(actualBody.Flags), "unexpected number of flags in response")
}

func TestGetFlag(t *testing.T) {
	flagRepo := fflag.NewRepo(testDB)

	savedFlag := fflag.Flag{
		ID:      "james-flag-3",
		Enabled: true,
	}
	err := flagRepo.Save(savedFlag)
	require.NoError(t, err, "could not save test data to DB")

	req := httptest.NewRequest(http.MethodGet, "/flags/james-flag-3", nil)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	resultBody := result.Body
	defer resultBody.Close()
	var actualBody fflag.Flag
	err = json.NewDecoder(resultBody).Decode(&actualBody)
	require.NoError(t, err, "could not decode response body")

	assert.Equal(t, savedFlag.ID, actualBody.ID)
	assert.Equal(t, savedFlag.Enabled, actualBody.Enabled)
}

func TestPostFlag(t *testing.T) {
	flagName := "my-flag"
	reqBody := fflag.PostReqBody{
		ID: flagName,
	}
	jsonReqBody, err := json.Marshal(reqBody)
	require.NoError(t, err, "could not marshal request body")

	req := httptest.NewRequest(http.MethodPost, "/flags", bytes.NewReader(jsonReqBody))
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	resultBody := result.Body
	defer resultBody.Close()
	var actualBody fflag.Flag
	err = json.NewDecoder(resultBody).Decode(&actualBody)
	require.NoError(t, err, "could not decode response body")

	assert.Equal(t, flagName, actualBody.ID)
	assert.Equal(t, false, actualBody.Enabled)
}

func TestPostFlag_invalidBody(t *testing.T) {
	flagName := "invalid flag name"
	reqBody := fflag.PostReqBody{
		ID: flagName,
	}
	jsonReqBody, err := json.Marshal(reqBody)
	require.NoError(t, err, "could not marshal request body")

	req := httptest.NewRequest(http.MethodPost, "/flags", bytes.NewReader(jsonReqBody))
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusBadRequest, result.StatusCode)
}

func TestDeleteFlag(t *testing.T) {
	flagRepo := fflag.NewRepo(testDB)

	flagName := "vans-flag"
	savedFlag := fflag.Flag{
		ID:      flagName,
		Enabled: true,
	}
	err := flagRepo.Save(savedFlag)
	require.NoError(t, err, "could not save test data to DB")

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/flags/%s", flagName), nil)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusNoContent, result.StatusCode)
}
