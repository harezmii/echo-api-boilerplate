package rest

import (
	_ "api/docs"
	"api/ent"
	"api/internal/handle"
	"api/internal/middleware"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
)

const Version = "0.0.5"

func SetupRest(address string) {
	app := echo.New()

	ent.ConnectionEnt()
	// Middleware
	middleware.SetupMiddleware(app)
	// Middleware END

	v1 := app.Group("/api/v1")
	// HANDLE SETUP
	handle.SetupHandle(v1)
	// HANDLE SETUP END

	// SWAGGER
	app.GET("/swagger/*", echoSwagger.WrapHandler)
	// SWAGGER END

	// Rest PING
	app.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, ".")
	})
	// REST PING END

	serverError := app.Start(":4000")
	if serverError != nil {
		return
	}
}
