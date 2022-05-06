package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/lkeix/graphql-sandbox/db"
	"github.com/lkeix/graphql-sandbox/graph"
	"github.com/lkeix/graphql-sandbox/graph/generated"
)

const (
	PORT = ":8080"
)

func main() {
	db.Connect()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello", hello)

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: db.DB}},
		),
	)

	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(ctx echo.Context) error {
		graphqlHandler.ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	})

	e.GET("/playground", func(ctx echo.Context) error {
		playgroundHandler.ServeHTTP(ctx.Response(), ctx.Request())
		return nil
	})

	e.Start(PORT)
}

func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "hello")
}

/*
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/lkeix/graphql-sandbox/graph"
	"github.com/lkeix/graphql-sandbox/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
*/
