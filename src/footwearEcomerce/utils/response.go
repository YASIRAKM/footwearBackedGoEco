package utils

import (
	"GoProjects/src/footweearEcomerce/src/footwearEcomerce/model"

	"github.com/labstack/echo/v4"
)

func Response(c echo.Context, httpStatus int, status bool, message string, data interface{}) error {
	response := model.ApiResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return c.JSON(httpStatus, response)
}
