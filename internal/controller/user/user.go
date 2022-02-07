package user

import (
	"api/internal/controller"
	"api/internal/model"
	"api/internal/response"
	"api/internal/secret"
	"api/internal/validation"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ControllerUser struct {
	controller.Controller
}

// Store ShowAccount godoc
// @Summary      Create Data
// @Description  create users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body body  model.User  false   "User form"
// @Success      201  {object}  []model.User
// @Router       /users [post]
func (u ControllerUser) Store(ctx echo.Context) error {
	// MODEL CREATE
	user := u.Model.(model.User)
	// MODEL CREATE END

	// BIND
	bindError := ctx.Bind(&user)
	if bindError != nil {
		fmt.Println("Bind Error")
	}
	// BIND END

	// VALIDATION
	validationError := validation.ValidateStructToTurkish(&user)
	if validationError != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{StatusCode: http.StatusUnprocessableEntity, Message: validationError})
	}
	// VALIDATION END

	// PASSWORD HASH
	hashPassword := secret.HashPassword(user.Password)
	// PASSWORD HASH END

	// DATABASE CREATE
	databaseError := u.Client.User.Create().SetStatus(*user.Status).SetEmail(user.Email).SetName(user.Name).SetPassword(hashPassword).Exec(u.Context)
	fmt.Println(databaseError.Error())
	if databaseError != nil {
		return ctx.JSON(http.StatusNoContent, response.ErrorResponse{StatusCode: http.StatusNoContent, Message: "Database create error"})
	}
	// DATABASE CREATE END

	user.Password = hashPassword
	return ctx.JSON(http.StatusCreated, response.SuccessResponse{StatusCode: http.StatusCreated, Message: "User created", Data: &user})
}
