package api2

import (
	"database/sql"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/j-dumbell/lite-flag/internal/auth"
	"github.com/j-dumbell/lite-flag/internal/fflag"
	"github.com/j-dumbell/lite-flag/internal/oapi"
)

var requestTimeout = 20 * time.Second

func toFlag(flagDTO oapi.Flag) (fflag.Flag, error) {
	flag := fflag.Flag{
		Key:      flagDTO.Key,
		Type:     fflag.FlagType(flagDTO.Type),
		IsPublic: flagDTO.IsPublic,
	}

	if flagDTO.Type == oapi.FlagTypeString {
		flagValue, err := flagDTO.Value.AsFlagValue1()
		if err != nil {
			return fflag.Flag{}, err
		}
		flag.StringValue = &flagValue
	}

	if flagDTO.Type == oapi.FlagTypeBoolean {
		flagValue, err := flagDTO.Value.AsFlagValue0()
		if err != nil {
			return fflag.Flag{}, err
		}
		flag.BooleanValue = &flagValue
	}

	// ToDo - json flag

	return flag, nil
}

func toFlagDTO(flag fflag.Flag) oapi.Flag {
	flagDTO := oapi.Flag{
		Key:      flag.Key,
		Type:     oapi.FlagType(flag.Type),
		IsPublic: flag.IsPublic,
	}

	switch flag.Type {
	case fflag.FlagTypeString:
		_ = flagDTO.Value.FromFlagValue1(*flag.StringValue)
	case fflag.FlagTypeBoolean:
		_ = flagDTO.Value.FromFlagValue0(*flag.BooleanValue)
	case fflag.FlagTypeJSON:
		// ToDo
	}

	return flagDTO
}

type server struct {
	db          *sql.DB
	flagService fflag.Service
	authService auth.Service
}

func New(db *sql.DB, flagService fflag.Service, authService auth.Service) chi.Router {
	srv := server{
		db:          db,
		flagService: flagService,
		authService: authService,
	}

	wrapper := oapi.ServerInterfaceWrapper{
		Handler: oapi.NewStrictHandler(&srv, nil),
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(requestTimeout))
	r.Use(newRoleMW(authService))

	r.Get("/flags", wrapper.GetFlags)
	r.With(adminOnly).Post("/flags", wrapper.PostFlags)

	r.With(adminOnly).Delete("/flags/{key}", wrapper.DeleteFlagsKey)
	r.With(adminOnly).Put("/flags/{key}", wrapper.PutFlagsKey)
	r.Get("/flags/{key}", wrapper.GetFlagsKey)

	return r
}
