package controller

import (
	"api/ent"
	"context"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	Client  *ent.Client
	Context context.Context
	Model   interface{}
}

type Control interface {
	Index(ctx echo.Context) error
	Store(ctx echo.Context) error
	Update(ctx echo.Context) error
	Destroy(ctx echo.Context) error
	Show(ctx echo.Context) error
}
