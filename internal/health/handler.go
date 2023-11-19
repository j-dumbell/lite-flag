package health

import (
	"database/sql"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) Handler {
	return Handler{db: db}
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := h.db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to connect to DB"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
	return
}
