package fflag

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type IRepo interface {
	Save(flag Flag) (Flag, error)
	FindOneByName(name string) (Flag, error)
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

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
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

		_, err := h.repo.FindOneByName(body.Name)
		if err == nil {
			w.WriteHeader(http.StatusConflict)
			return
		}
		if err != sql.ErrNoRows {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		flag, err := h.repo.Save(Flag{
			Name:      body.Name,
			CreatedAt: time.Now(),
		})
		if err != nil {
			fmt.Println(err)
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
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
