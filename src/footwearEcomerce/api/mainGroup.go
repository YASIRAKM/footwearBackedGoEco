package api

import (
	"GoProjects/src/footweearEcomerce/src/footwearEcomerce/api/handler"

	"github.com/labstack/echo/v4"
)

func MainGroup(e *echo.Echo) {

	e.POST("/login", handler.Login)
	e.POST("/getuser", handler.GetUserByUsername)
}
