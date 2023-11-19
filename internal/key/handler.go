package key

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/j-dumbell/lite-flag/pkg/array"
)

type iRepo interface {
	Save(apiKey ApiKey) error
	FindAll() ([]ApiKey, error)
	DeleteByID(id string) error
	FindOneByID(id string) (ApiKey, error)
	FindOneByKey(key string) (ApiKey, error)
	FindOne(filters Filters) (ApiKey, error)
}

type Handler struct {
	repo iRepo
}

func NewHandler(repo iRepo) Handler {
	return Handler{repo: repo}
}

type PostReqBody struct {
	ID   string `json:"id"`
	Role Role   `json:"role"`
}

// ToDo - collect all errors?
func (body PostReqBody) Validate() error {
	if body.ID == "" {
		return errors.New("name must be provided")
	}
	if body.Role == "" {
		return errors.New("role must be provided")
	}
	if !body.Role.isValid() {
		return errors.New("role is invalid")
	}
	return nil
}

func newKey() string {
	b := make([]byte, 40)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (h Handler) Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var postApiKeyBody PostReqBody
	if err := json.NewDecoder(r.Body).Decode(&postApiKeyBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := postApiKeyBody.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// ToDo json structure
		w.Write([]byte(err.Error()))
		return
	}

	_, err := h.repo.FindOneByID(postApiKeyBody.ID)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		bytes, err := json.Marshal(errors.New("an API key with that id already exists"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Write(bytes)
		return
	}
	if err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	apiKeyModel := ApiKey{ID: postApiKeyBody.ID, ApiKey: newKey(), CreatedAt: time.Now(), Role: postApiKeyBody.Role}
	err = h.repo.Save(apiKeyModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(apiKeyModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)
}

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	keyID := chi.URLParam(r, "keyID")

	_, err := h.repo.FindOneByID(keyID)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			w.WriteHeader(http.StatusNotFound)
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	err = h.repo.DeleteByID(keyID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type GetApiKeysResponse struct {
	Keys []ApiKeyRedacted `json:"keys"`
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	apiKeys, err := h.repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	redactedKeys := array.ArrMap(apiKeys, RemoveKey)
	bytes, err := json.Marshal(GetApiKeysResponse{
		Keys: redactedKeys,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
