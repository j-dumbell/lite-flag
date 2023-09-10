package fflag

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type IRepo interface {
	Save(flag Flag) error
	FindOne(name string) (Flag, error)
	FindAll() ([]Flag, error)
}

type Handler struct {
	repo IRepo
}

func NewHandler(repo IRepo) Handler {
	return Handler{
		repo: repo,
	}
}

type postReqBody struct {
	Name string `json:"name"`
}

func (body postReqBody) validate() error {
	if strings.ContainsRune(body.Name, ' ') || strings.ContainsRune(body.Name, '/') {
		return errors.New("name can only contain letters, numbers, hyphens or underscores")
	}
	return nil
}

func (h Handler) post(w http.ResponseWriter, r *http.Request) {
	var body postReqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := body.validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	_, err := h.repo.FindOne(body.Name)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	flag := Flag{
		Name:      body.Name,
		Enabled:   false,
		CreatedAt: time.Now(),
		Schedule:  nil,
	}

	err = h.repo.Save(flag)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	flagResponse, err := json.Marshal(flag)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(flagResponse)
}

type GetResponse struct {
	Flags []Flag `json:"flags"`
}

func (h Handler) getAll(w http.ResponseWriter, r *http.Request) {
	flags, err := h.repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(flags)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h Handler) getOne(w http.ResponseWriter, r *http.Request) {
	name := namePathRegexp.FindStringSubmatch(r.URL.Path)[1]

	flag, err := h.repo.FindOne(name)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(flag)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

var namePathRegexp = regexp.MustCompile(`^\/flags\/([A-z0-9-_]+)$`)

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if namePathRegexp.MatchString(r.URL.Path) {
		switch r.Method {
		case http.MethodGet:
			h.getOne(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
		return
	} else {
		switch r.Method {
		case http.MethodPost:
			h.post(w, r)
		case http.MethodGet:
			h.getAll(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}
