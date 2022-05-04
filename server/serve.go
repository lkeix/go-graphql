package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lkeix/graphql-sandbox/db"
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

	e.Start(PORT)
}

func hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "hello")
}
