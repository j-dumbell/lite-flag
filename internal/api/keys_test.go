package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/j-dumbell/lite-flag/internal/key"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetKeys(t *testing.T) {
	keyRepo := key.NewRepo(testDB, cipherKey)

	key1 := key.ApiKey{
		ID:     "key-1",
		ApiKey: "adsfsj;klasdjf13asdj;123kjas;d123asdk21aslkd123123kasd",
		Role:   key.Admin,
	}
	key2 := key.ApiKey{
		ID:     "key-2",
		ApiKey: "adsfsj;klasdjf13asdasdj123jalsdkjas;d123asdk21aslkd123123kasd",
		Role:   key.Root,
	}
	err1 := keyRepo.Save(key1)
	err2 := keyRepo.Save(key2)
	require.NoError(t, err1, "could not save test data to DB")
	require.NoError(t, err2, "could not save test data to DB")

	req := httptest.NewRequest(http.MethodGet, "/api-keys", nil)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusOK, result.StatusCode)

	resultBody := result.Body
	defer resultBody.Close()
	var actualBody key.GetApiKeysResponse
	err := json.NewDecoder(resultBody).Decode(&actualBody)
	require.NoError(t, err, "could not decode response body")

	assert.Equal(t, 2, len(actualBody.Keys), "unexpected number of keys in response")
	expectedKeys := []key.ApiKeyRedacted{
		{
			ID:   key1.ID,
			Role: key1.Role,
		},
		{
			ID:   key2.ID,
			Role: key2.Role,
		},
	}
	assert.Subset(t, expectedKeys, actualBody.Keys)
}

func TestPostKey(t *testing.T) {
	keyName := "some-api-key"
	reqBody := key.PostReqBody{
		ID:   keyName,
		Role: key.Admin,
	}
	jsonReqBody, err := json.Marshal(reqBody)
	require.NoError(t, err, "could not marshal request body")

	req := httptest.NewRequest(http.MethodPost, "/api-keys", bytes.NewReader(jsonReqBody))
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusCreated, result.StatusCode)

	resultBody := result.Body
	defer resultBody.Close()

	var actualBody key.ApiKey
	err = json.NewDecoder(resultBody).Decode(&actualBody)
	require.NoError(t, err, "could not decode response body")

	assert.Equal(t, reqBody.ID, actualBody.ID)
	assert.Equal(t, reqBody.Role, actualBody.Role)
	assert.NotEmpty(t, actualBody.ApiKey)
	assert.NotEmpty(t, actualBody.ApiKey)
}

func TestPostKey_InvalidReq(t *testing.T) {
	keyName := "blah-3"
	reqBody := key.PostReqBody{
		ID:   keyName,
		Role: "",
	}
	jsonReqBody, err := json.Marshal(reqBody)
	require.NoError(t, err, "could not marshal request body")

	req := httptest.NewRequest(http.MethodPost, "/api-keys", bytes.NewReader(jsonReqBody))
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusBadRequest, result.StatusCode)
}

func TestDeleteKey(t *testing.T) {
	keyRepo := key.NewRepo(testDB, cipherKey)

	keyName := "blah-key"
	err := keyRepo.Save(key.ApiKey{
		ID:     keyName,
		ApiKey: "asdaskld21312klasdklj1kl2j3jjsadlkjjq2klj3lk12jkdjsal",
		Role:   key.Admin,
	})
	require.NoError(t, err, "could not insert test data")

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api-keys/%s", keyName), nil)
	w := httptest.NewRecorder()
	api.ServeHTTP(w, req)

	result := w.Result()
	assert.Equal(t, http.StatusNoContent, result.StatusCode)
}
