package key

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
)

type iRepo interface {
	Save(apiKey ApiKey) (ApiKey, error)
	NameExists(name string) (bool, error)
	FindAll() ([]ApiKey, error)
	DeleteById(id int) error
	IdExists(id int) (bool, error)
}

type Handler struct {
	repo iRepo
}

func NewHandler(db *sql.DB) Handler {
	return Handler{repo: NewRepo(db)}
}

type postReqBody struct {
	Name string `json:"name"`
}

var idPathRegexp = regexp.MustCompile(`^\/api-keys\/(\d+)$`)

func newKey() string {
	b := make([]byte, 40)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	switch {
	case r.Method == http.MethodPost:
		var postApiKeyBody postReqBody
		if err := json.NewDecoder(r.Body).Decode(&postApiKeyBody); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			//ToDo - return error message in body.  Validation library?
			return
		}

		exists, err := h.repo.NameExists(postApiKeyBody.Name)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if exists {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("an API key with that name already exists"))
			return
		}

		apiKeyModel := New(postApiKeyBody.Name)

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

	case r.Method == http.MethodDelete && idPathRegexp.MatchString(r.URL.Path):
		apiKeyId, _ := strconv.Atoi(idPathRegexp.FindStringSubmatch(r.URL.Path)[1])
		exists, err := h.repo.IdExists(apiKeyId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if exists == false {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		err = h.repo.DeleteById(apiKeyId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
