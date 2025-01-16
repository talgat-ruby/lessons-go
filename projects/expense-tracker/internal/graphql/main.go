package graphql

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/graphql/graph"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/graphql/graph/directives"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/graphql/graph/middleware"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker/internal/types/controller"
)

type GraphQL struct {
	conf   *config.APIGraphQLConfig
	logger *slog.Logger
	ctrl   controller.Controller
}

func New(conf *config.APIGraphQLConfig, logger *slog.Logger, ctrl controller.Controller) *GraphQL {
	return &GraphQL{
		conf:   conf,
		logger: logger,
		ctrl:   ctrl,
	}
}

func (a *GraphQL) Start(ctx context.Context) error {
	midd := middleware.New(a.logger, a.ctrl)
	d := directives.New(a.logger, a.ctrl)

	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: graph.NewResolver(a.conf, a.ctrl, a.logger),
				Directives: graph.DirectiveRoot{
					Auth: d.AuthDirective,
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", midd.Authenticator(srv))

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", a.conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", a.conf.Port), nil))

	return nil
}
