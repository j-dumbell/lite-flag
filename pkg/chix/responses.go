package chix

import (
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func NewMessage(message string) Message {
	return Message{message}
}

type Response struct {
	StatusCode int
	Body       any
}

func OK(body any) Response {
	return Response{
		StatusCode: http.StatusOK,
		Body:       body,
	}
}

func NotFound(body any) Response {
	return Response{
		StatusCode: http.StatusNotFound,
		Body:       body,
	}
}

func Conflict(body any) Response {
	return Response{
		StatusCode: http.StatusConflict,
		Body:       body,
	}
}

func InternalServerError() Response {
	return Response{
		StatusCode: http.StatusInternalServerError,
		Body:       nil,
	}
}

func MethodNotAllowed() Response {
	return Response{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       nil,
	}
}

func NoContent() Response {
	return Response{
		StatusCode: http.StatusNoContent,
		Body:       nil,
	}
}

func BadRequest(body any) Response {
	return Response{
		StatusCode: http.StatusBadRequest,
		Body:       body,
	}
}

func Created(body any) Response {
	return Response{
		StatusCode: http.StatusCreated,
		Body:       body,
	}
}
