package rest

import (
	"api/internal/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupRest(port int) {
	app := echo.New()

	// Middleware
	middleware.SetupMiddleware(app)
	// Middleware END

	// Rest PING
	app.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, ".")
	})
	// REST PING END

	app.Logger.Fatal(app.Start(":4000"))
}
