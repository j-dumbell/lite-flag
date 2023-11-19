package fflag

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

type IRepo interface {
	Save(flag Flag) error
	FindOne(id string) (Flag, error)
	FindAll() ([]Flag, error)
	Delete(id string) error
}

type Handler struct {
	repo IRepo
}

func NewHandler(repo IRepo) Handler {
	return Handler{
		repo: repo,
	}
}

type PostReqBody struct {
	ID string `json:"id"`
}

func (body PostReqBody) Validate() error {
	if strings.ContainsRune(body.ID, ' ') || strings.ContainsRune(body.ID, '/') {
		return errors.New("name can only contain letters, numbers, hyphens or underscores")
	}
	return nil
}

func (h Handler) Post(w http.ResponseWriter, r *http.Request) {
	var body PostReqBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := body.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	_, err := h.repo.FindOne(body.ID)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	if err != sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	flag := Flag{
		ID:        body.ID,
		Enabled:   false,
		CreatedAt: time.Now(),
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

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	flags, err := h.repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(GetResponse{Flags: flags})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (h Handler) GetOne(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "flagID")

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

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "flagID")

	_, err := h.repo.FindOne(name)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.repo.Delete(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
