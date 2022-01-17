package main

import (
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sgoodliff/gogogo/internal/users"
)

var (
	port int
)

/*
// User
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}
*/
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
		state, _ := users.IsValidUser(username, password)
		if state == true {
			return true, nil
		}
		return false, nil

	}))

	e.GET("/users/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/:id")
	})

	e.GET("/users/new", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/new")
	})
	e.GET("/users/list", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/list")
	})

	e.GET("/users/:id/json", func(c echo.Context) error {
		id := c.Param("id")
		u := users.Lookup(id)
		return c.JSONPretty(http.StatusOK, u, "  ")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
