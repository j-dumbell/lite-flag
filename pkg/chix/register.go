package chix

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler func(*http.Request) Response

func writeResponse(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := handler(r)

		if response.Body == nil {
			w.WriteHeader(response.StatusCode)
			return
		}

		var toMarshal any
		switch t := response.Body.(type) {
		case string:
			toMarshal = MessageResponse{t}
		default:
			toMarshal = t
		}

		jsonBody, err := json.Marshal(toMarshal)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		_, _ = w.Write(jsonBody)
	}
}

func Get(router *chi.Mux, pattern string, handler Handler, middlewares ...func(http.Handler) http.Handler) {
	router.With(middlewares...).Get(pattern, writeResponse(handler))
}

func Post(router *chi.Mux, pattern string, handler Handler, middlewares ...func(http.Handler) http.Handler) {
	router.With(middlewares...).Post(pattern, writeResponse(handler))
}

func Put(router *chi.Mux, pattern string, handler Handler, middlewares ...func(http.Handler) http.Handler) {
	router.With(middlewares...).Put(pattern, writeResponse(handler))
}

func Patch(router *chi.Mux, pattern string, handler Handler, middlewares ...func(http.Handler) http.Handler) {
	router.With(middlewares...).Patch(pattern, writeResponse(handler))
}

func Delete(router *chi.Mux, pattern string, handler Handler, middlewares ...func(http.Handler) http.Handler) {
	router.With(middlewares...).Delete(pattern, writeResponse(handler))
}
