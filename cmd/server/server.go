package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	port int
)

func main() {
	e := echo.New()
	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	// Middleware
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		/*
			if username == "joe" && password == "secret" {
				return true, nil
			}
			return false, nil
		*/
		users.i
	}))

	e.Logger.Fatal(e.Start(":8080"))
}
