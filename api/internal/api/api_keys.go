package api

import (
	"context"
	"errors"

	"github.com/j-dumbell/lite-flag/internal/auth"
	"github.com/j-dumbell/lite-flag/internal/oapi"
	"github.com/j-dumbell/lite-flag/pkg/validation"
)

func (srv *server) PostApiKeys(ctx context.Context, request oapi.PostApiKeysRequestObject) (oapi.PostApiKeysResponseObject, error) {
	user, _ := getUser(ctx)
	if user.Role == auth.RoleAdmin && request.Body.Role != oapi.ApiKeyInputRoleReadonly {
		return oapi.PostApiKeys403JSONResponse(map[string]interface{}{"error": "admins may only create API keys with readonly role"}), nil
	}

	createApiKeyParams := auth.CreateApiKeyParams{
		Name: request.Body.Name,
		Role: auth.Role(request.Body.Role),
	}

	apiKey, err := srv.authService.CreateKey(ctx, createApiKeyParams)
	if errors.As(err, &validation.Result{}) {
		return oapi.PostApiKeys400JSONResponse(map[string]interface{}{"error": err}), nil
	} else if errors.Is(err, auth.ErrAlreadyExists) {
		return oapi.PostApiKeys409Response{}, nil
	} else if err != nil {
		return nil, err
	}

	return oapi.PostApiKeys201JSONResponse(toApiKeyDTO(apiKey)), nil
}

func (srv *server) DeleteApiKeysName(ctx context.Context, request oapi.DeleteApiKeysNameRequestObject) (oapi.DeleteApiKeysNameResponseObject, error) {
	apiKeyRedacted, err := srv.authService.FindOneByName(ctx, request.Name)
	if errors.Is(err, auth.ErrKeyNotFound) {
		return oapi.DeleteApiKeysName404Response{}, nil
	} else if err != nil {
		return nil, err
	}

	requestor, _ := getUser(ctx)
	if apiKeyRedacted.Role == auth.RoleRoot {
		return oapi.DeleteApiKeysName403JSONResponse(map[string]interface{}{"error": "root API key cannot be deleted"}), nil
	}
	if requestor.Role == auth.RoleAdmin && apiKeyRedacted.Role != auth.RoleReadonly {
		return oapi.DeleteApiKeysName403JSONResponse(map[string]interface{}{"error": "admins may only delete readonly API keys"}), nil
	}

	if err := srv.authService.DeleteByName(ctx, request.Name); err != nil {
		return nil, err
	}

	return oapi.DeleteApiKeysName204Response{}, nil
}

func (srv *server) PostApiKeysNameRotate(ctx context.Context, request oapi.PostApiKeysNameRotateRequestObject) (oapi.PostApiKeysNameRotateResponseObject, error) {
	apiKeyRedacted, err := srv.authService.FindOneByName(ctx, request.Name)
	if errors.Is(err, auth.ErrKeyNotFound) {
		return oapi.PostApiKeysNameRotate404Response{}, nil
	} else if err != nil {
		return nil, err
	}

	user, _ := getUser(ctx)
	if user.Role == auth.RoleAdmin && apiKeyRedacted.Name != user.Name {
		return oapi.PostApiKeysNameRotate403JSONResponse(map[string]interface{}{"error": "admins may only rotate their own keys or readonly keys"}), nil
	}

	newApiKey, err := srv.authService.RotateKey(ctx, request.Name)
	if err != nil {
		return nil, err
	}

	return oapi.PostApiKeysNameRotate201JSONResponse(toApiKeyDTO(newApiKey)), nil
}
