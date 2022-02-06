package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupMiddleware(echo *echo.Echo) {
	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())
}
