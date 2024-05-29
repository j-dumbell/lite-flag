package chix

import (
	"net/http"
)

type MessageResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error any `json:"error"`
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

func Forbidden(body any) Response {
	return Response{
		StatusCode: http.StatusForbidden,
		Body:       body,
	}
}

func Unauthorized(body any) Response {
	return Response{
		StatusCode: http.StatusUnauthorized,
		Body:       body,
	}
}
