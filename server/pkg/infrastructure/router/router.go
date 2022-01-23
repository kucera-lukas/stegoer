package router

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"

	"github.com/kucera-lukas/stegoer/ent"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/env"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/log"
	"github.com/kucera-lukas/stegoer/pkg/infrastructure/middleware"
)

// Routes of Router.
const (
	QueryPath      = "/graphql"
	PlaygroundPath = "/playground"
)

// New creates new mux router.
func New(
	config *env.Config,
	logger *log.Logger,
	srv http.Handler,
	client *ent.Client,
) *mux.Router {
	router := mux.NewRouter()
	router.Use(middleware.Logging, middleware.Jwt(logger, client))

	router.Handle(QueryPath, srv)

	if config.Debug {
		router.HandleFunc(
			PlaygroundPath,
			playground.Handler("GQL Playground", QueryPath),
		)
	}

	return router
}
