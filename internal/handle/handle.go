package handle

import (
	"api/ent"
	"api/internal/controller/user"
	"api/internal/model"
	"context"
	"github.com/labstack/echo/v4"
)

var connection = ent.ConnectionEnt()
var contextBack = context.Background()

func SetupHandle(echo *echo.Group) {
	userController := user.ControllerUser{
		Controller: struct {
			Client  *ent.Client
			Context context.Context
			Model   interface{}
		}{Client: connection, Context: contextBack, Model: model.User{}},
	}
	echo.POST("/users", userController.Store)
	echo.GET("/users/:id", userController.Show)
	echo.DELETE("/users/:id", userController.Destroy)
	echo.PUT("/users/:id", userController.Update)

}
