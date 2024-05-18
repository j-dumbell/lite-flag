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

		body := response.Body
		message, isString := body.(string)
		if isString {
			body = NewMessage(message)
		}

		jsonBody, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(response.StatusCode)
		w.Write(jsonBody)
	}
}

func RegisterGet(router *chi.Mux, pattern string, handler Handler) {
	router.Get(pattern, writeResponse(handler))
}

func RegisterPost(router *chi.Mux, pattern string, handler Handler) {
	router.Post(pattern, writeResponse(handler))
}

func RegisterPut(router *chi.Mux, pattern string, handler Handler) {
	router.Put(pattern, writeResponse(handler))
}

func RegisterPatch(router *chi.Mux, pattern string, handler Handler) {
	router.Patch(pattern, writeResponse(handler))
}

func RegisterDelete(router *chi.Mux, pattern string, handler Handler) {
	router.Delete(pattern, writeResponse(handler))
}
