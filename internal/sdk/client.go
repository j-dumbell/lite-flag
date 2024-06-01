package sdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/j-dumbell/lite-flag/internal/api"
	"github.com/j-dumbell/lite-flag/internal/fflag"
)

type Client struct {
	host   string
	apiKey *string
}

func NewClient(host string, apiKey *string) Client {
	return Client{
		host:   host,
		apiKey: apiKey,
	}
}

var ErrNoAPIKeyProvided = errors.New("no API key provided")
var ErrUnauthorized = errors.New("unauthorized")

func (client *Client) GetFlags(ctx context.Context) ([]fflag.Flag, error) {
	return get[[]fflag.Flag](ctx, *client, "flags")
}

func (client *Client) GetFlag(ctx context.Context, key string) (fflag.Flag, error) {
	return get[fflag.Flag](ctx, *client, fmt.Sprintf("flag/%s", key))
}

var ErrNotStringFlag = errors.New("flag is not a string flag")
var ErrNotBooleanFlag = errors.New("flag is not a boolean flag")
var ErrNotJSONFlag = errors.New("flag is not a JSON flag")
var ErrFlagValueNil = errors.New("flag value is nil")

func (client *Client) GetStringFlagValue(ctx context.Context, key string) (string, error) {
	flag, err := client.GetFlag(ctx, key)
	if err != nil {
		return "", err
	}

	if flag.Type != fflag.FlagTypeString {
		return "", ErrNotStringFlag
	}

	if flag.StringValue == nil {
		return "", ErrFlagValueNil
	}

	return *flag.StringValue, nil
}

func (client *Client) GetBooleanFlagValue(ctx context.Context, key string) (bool, error) {
	flag, err := client.GetFlag(ctx, key)
	if err != nil {
		return false, err
	}

	if flag.Type != fflag.FlagTypeBoolean {
		return false, ErrNotBooleanFlag
	}

	if flag.BooleanValue == nil {
		return false, ErrFlagValueNil
	}

	return *flag.BooleanValue, nil
}

func (client *Client) GetJSONFlagValue(ctx context.Context, key string, v any) error {
	flag, err := client.GetFlag(ctx, key)
	if err != nil {
		return err
	}

	if flag.Type != fflag.FlagTypeJSON {
		return ErrNotStringFlag
	}

	jsonBytes, err := json.Marshal(flag.JSONValue)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonBytes, v); err != nil {
		return err
	}

	return nil
}

var ErrUnrecognizedStatusCode = errors.New("unrecognized status code")

func get[T any](ctx context.Context, client Client, endpoint string) (T, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/%s", client.host, endpoint), nil)
	if err != nil {
		var t T
		return t, fmt.Errorf("failed to initialize request: %w", err)
	}

	if client.apiKey != nil {
		req.Header.Set(api.ApiKeyHeader, *client.apiKey)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		var t T
		return t, fmt.Errorf("failed to execute request: %w", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
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
