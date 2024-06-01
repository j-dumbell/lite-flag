package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/j-dumbell/lite-flag/internal/api"
)

var ErrUnrecognizedStatusCode = errors.New("unrecognized status code")

func buildURL(parts ...string) string {
	trimmedParts := []string{}

	for i, part := range parts {
		var trimmedPart string
		if i == 0 {
			trimmedPart = strings.TrimRight(part, "/")
		} else {
			trimmedPart = strings.Trim(part, "/")
		}

		if trimmedPart != "" {
			trimmedParts = append(trimmedParts, trimmedPart)
		}
	}

	return strings.Join(trimmedParts, "/")
}

func request[T any](ctx context.Context, client Client, method string, endpoint string, body any) (T, error) {
	reqBodyBytes, err := json.Marshal(body)
	if err != nil {
		var t T
		return t, fmt.Errorf("failed to marshal body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, method, buildURL(client.host, endpoint), bytes.NewBuffer(reqBodyBytes))
	if err != nil {
		var t T
		return t, fmt.Errorf("failed to initialize request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	if client.apiKey != nil {
		req.Header.Set(api.ApiKeyHeader, *client.apiKey)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		var t T
		return t, fmt.Errorf("failed to execute request: %w", err)
	}

	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		var t T

		msg, exists := statusCodes[resp.StatusCode]
		if !exists {
			return t, ErrUnrecognizedStatusCode
		}

		return t, errors.New(msg)
	}

	var t T
	err = json.NewDecoder(resp.Body).Decode(&t)
	if err != nil {
		return t, fmt.Errorf("failed to decode response body: %w", err)
	}

	return t, nil
}

func get[T any](ctx context.Context, client Client, endpoint string) (T, error) {
	return request[T](ctx, client, http.MethodGet, endpoint, nil)
}

func post[T any](ctx context.Context, client Client, endpoint string, body any) (T, error) {
	return request[T](ctx, client, http.MethodPost, endpoint, body)
}

func put[T any](ctx context.Context, client Client, endpoint string, body any) (T, error) {
	return request[T](ctx, client, http.MethodPut, endpoint, body)
}

func delete[T any](ctx context.Context, client Client, endpoint string) (T, error) {
	return request[T](ctx, client, http.MethodDelete, endpoint, nil)
}
