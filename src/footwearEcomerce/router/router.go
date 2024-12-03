package router

import (
	"GoProjects/src/footweearEcomerce/src/footwearEcomerce/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	// Define a simple route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	api.MainGroup(e)

	return e

}
