package key

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"github.com/j-dumbell/lite-flag/pkg/array"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type iRepo interface {
	Save(apiKey ApiKey) (ApiKey, error)
	FindAll() ([]ApiKey, error)
	DeleteById(id int) error
	FindOneById(id int) (ApiKey, error)
	FindOneByName(name string) (ApiKey, error)
	FindOne(filters Filters) (ApiKey, error)
}

type Handler struct {
	repo iRepo
}

func NewHandler(db *sql.DB) Handler {
	return Handler{repo: NewRepo(db)}
}

type postReqBody struct {
	Name string `json:"name"`
	Role Role   `json:"role"`
}

func (body postReqBody) isValid() bool {
	return body.Name != "" && body.Role.isValid()
}

var idPathRegexp = regexp.MustCompile(`^\/api-keys\/(\d+)$`)

func newKey() string {
	b := make([]byte, 40)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (h Handler) post(w http.ResponseWriter, r *http.Request) {
	var postApiKeyBody postReqBody
	if err := json.NewDecoder(r.Body).Decode(&postApiKeyBody); err != nil || !postApiKeyBody.isValid() {
		w.WriteHeader(http.StatusBadRequest)
		//ToDo - return error message in body.  Validation library?
		return
	}

	_, err := h.repo.FindOneByName(postApiKeyBody.Name)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("an API key with that name already exists"))
		return
	}
	if err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	apiKeyModel := ApiKey{Name: postApiKeyBody.Name, ApiKey: newKey(), CreatedAt: time.Now(), Role: postApiKeyBody.Role}

	apiKeyModel, err = h.repo.Save(apiKeyModel)
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

func (h Handler) delete(w http.ResponseWriter, r *http.Request) {
	apiKeyId, _ := strconv.Atoi(idPathRegexp.FindStringSubmatch(r.URL.Path)[1])
	existingApiKey, err := h.repo.FindOneById(apiKeyId)
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
	if existingApiKey.Role == Root {
		//ToDo check this is the correct status code
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("cannot delete root API key"))
		return
	}

	err = h.repo.DeleteById(apiKeyId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h Handler) get(w http.ResponseWriter, r *http.Request) {
	apiKeys, err := h.repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	redactedKeys := array.ArrMap(apiKeys, RemoveKey)
	bytes, err := json.Marshal(redactedKeys)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	switch {
	case r.Method == http.MethodPost:
		h.post(w, r)

	case r.Method == http.MethodDelete && idPathRegexp.MatchString(r.URL.Path):
		h.delete(w, r)

	case r.Method == http.MethodGet && !idPathRegexp.MatchString(r.URL.Path):
		h.get(w, r)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
