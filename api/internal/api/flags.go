package api

import (
	"context"
	"errors"

	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/internal/oapi"
	"github.com/j-dumbell/lite-flag/pkg/fp"
	"github.com/j-dumbell/lite-flag/pkg/validation"
)

func (srv *server) GetFlags(ctx context.Context, _ oapi.GetFlagsRequestObject) (oapi.GetFlagsResponseObject, error) {
	_, authenticated := getUser(ctx)

	flags, err := srv.flagService.FindAll(ctx, !authenticated)
	if err != nil {
		return nil, err
	}

	return oapi.GetFlags200JSONResponse(fp.Map(flags, toFlagDTO)), nil
}

func (srv *server) GetFlagsKey(ctx context.Context, request oapi.GetFlagsKeyRequestObject) (oapi.GetFlagsKeyResponseObject, error) {
	flag, err := srv.flagService.FindOne(ctx, request.Key)
	if errors.Is(err, fflag.ErrNotFound) {
		return oapi.GetFlagsKey404Response{}, nil
	} else if err != nil {
		return nil, err
	}

	_, isAuthenticated := getUser(ctx)
	if !flag.IsPublic && !isAuthenticated {
		return oapi.GetFlagsKey403Response{}, nil
	}

	return oapi.GetFlagsKey200JSONResponse(toFlagDTO(flag)), nil
}

func (srv *server) PostFlags(ctx context.Context, request oapi.PostFlagsRequestObject) (oapi.PostFlagsResponseObject, error) {
	if request.Body == nil {
		return oapi.PostFlags400JSONResponse(map[string]interface{}{"error": "no body provided"}), nil
	}

	flag, err := toFlag(*request.Body)
	if err != nil {
		return oapi.PostFlags400JSONResponse(map[string]interface{}{"error": err.Error()}), nil
	}

	createdFlag, err := srv.flagService.Create(ctx, flag)
	if errors.Is(err, fflag.ErrAlreadyExists) {
		return oapi.PostFlags409Response{}, nil
	} else if err != nil {
		return nil, err
	}

	return oapi.PostFlags201JSONResponse(toFlagDTO(createdFlag)), nil
}

func (srv *server) DeleteFlagsKey(ctx context.Context, request oapi.DeleteFlagsKeyRequestObject) (oapi.DeleteFlagsKeyResponseObject, error) {
	err := srv.flagService.Delete(ctx, request.Key)
	if errors.Is(err, fflag.ErrNotFound) {
		return oapi.DeleteFlagsKey404Response{}, nil
	} else if err != nil {
		return nil, err
	}

	return oapi.DeleteFlagsKey204Response{}, nil
}

func (srv *server) PutFlagsKey(ctx context.Context, request oapi.PutFlagsKeyRequestObject) (oapi.PutFlagsKeyResponseObject, error) {
	flagDTO := oapi.Flag{
		Key:      request.Key,
		IsPublic: request.Body.IsPublic,
		Type:     oapi.FlagType(request.Body.Type),
		Value:    oapi.Flag_Value(request.Body.Value),
	}

	flag, err := toFlag(flagDTO)
	if err != nil {
		// ToDo - should we return a 400 here?
		return nil, err
	}

	updatedFlag, err := srv.flagService.Update(ctx, flag)
	switch {
	case errors.Is(err, fflag.ErrNotFound):
		return oapi.PutFlagsKey404Response{}, nil
	case errors.As(err, &validation.Result{}):
		return oapi.PutFlagsKey400JSONResponse(map[string]interface{}{"errors": err}), nil
	case err != nil:
		return nil, err
	}

	return oapi.PutFlagsKey200JSONResponse(toFlagDTO(updatedFlag)), nil
}
