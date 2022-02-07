package user

import (
	"api/internal/controller"
	"api/internal/dto"
	"api/internal/model"
	"api/internal/response"
	"api/internal/secret"
	"api/internal/validation"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"strings"
	"time"
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

// Update ShowAccount godoc
// @Summary      update Data
// @Description  update users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        body body  model.User  false   "User form"
// @Param        id  path  string  true   "User ID"
// @Success      200  {object}  model.User
// @Router       /users/{id} [put]
func (u ControllerUser) Update(ctx echo.Context) error {
	// MODEL CREATE
	user := u.Model.(model.User)
	// MODEL CREATE END

	// BIND
	bindError := ctx.Bind(&user)
	if bindError != nil {
		fmt.Println("Bind Error")
	}
	// BIND END

	// ID Convert
	id := ctx.Param("id")
	idInt, convertError := strconv.Atoi(id)
	if convertError != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Id convert error"})
	}
	// ID Convert END

	hashedPassword := secret.HashPassword(user.Password)
	validateError := validation.ValidateStructToTurkish(user)
	if validateError == nil {
		selectId, err := u.Client.User.Query().Where(func(s *sql.Selector) {
			s.Where(sql.IsNull("deleted_at"))
			s.Where(sql.EQ("id", idInt))
		}).FirstID(u.Context)

		if selectId != 0 {
			errt := u.Client.User.UpdateOneID(idInt).SetName(user.Name).SetPassword(hashedPassword).SetEmail(user.Email).SetStatus(*user.Status).SetUpdatedAt(time.Now()).Exec(u.Context)
			if errt != nil {
				return ctx.JSON(http.StatusNotFound, response.ErrorResponse{StatusCode: 404, Message: "user not updated, " + strings.Split(errt.Error(), ":")[3]})
			}
		}

		if err != nil {
			return ctx.JSON(http.StatusNotFound, response.ErrorResponse{StatusCode: 404, Message: "User not updated"})
		}

		return ctx.JSON(http.StatusOK, response.SuccessResponse{StatusCode: 200, Message: "User updated", Data: user})
	}
	return ctx.JSON(http.StatusUnprocessableEntity, response.ErrorResponse{StatusCode: http.StatusUnprocessableEntity, Message: validateError})
}

// Show ShowAccount godoc
// @Summary      Show Data
// @Description  get string by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "User ID"
// @Success      200  {object}  model.User
// @Router       /users/{id} [get]
func (u ControllerUser) Show(ctx echo.Context) error {
	// ID Convert
	id := ctx.Param("id")
	idInt, convertError := strconv.Atoi(id)
	if convertError != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Id convert error"})
	}
	// ID Convert END

	var responseDto []dto.UserDto
	getError := u.Client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).Select("id", "name", "email", "status").Scan(u.Context, &responseDto)
	if getError != nil {
		return ctx.JSON(http.StatusNoContent, response.ErrorResponse{StatusCode: http.StatusNoContent, Message: "Database error"})
	}

	// Deleted record find
	if len(responseDto) == 0 {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse{StatusCode: http.StatusNotFound, Message: "User not finding"})
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse{StatusCode: http.StatusOK, Message: "User find", Data: responseDto})
}

// Destroy ShowAccount godoc
// @Summary      Delete Data
// @Description  get string by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id  path  string  true   "User ID"
// @Success      200  {object}  model.User
// @Router       /users/{id} [delete]
func (u ControllerUser) Destroy(ctx echo.Context) error {
	// ID Convert
	id := ctx.Param("id")
	idInt, convertError := strconv.Atoi(id)
	if convertError != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse{StatusCode: http.StatusBadRequest, Message: "Id convert error"})
	}
	// ID Convert END

	firstID, queryError := u.Client.User.Query().Where(func(s *sql.Selector) {
		s.Where(sql.IsNull("deleted_at"))
		s.Where(sql.EQ("id", idInt))
	}).FirstID(u.Context)
	fmt.Println(firstID)
	if queryError != nil {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse{StatusCode: http.StatusNotFound, Message: "Database error"})
	}
	if firstID == 0 {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse{StatusCode: http.StatusNotFound, Message: "Database error"})
	}
	databaseError := u.Client.User.UpdateOneID(firstID).SetDeletedAt(time.Now()).Exec(u.Context)
	if databaseError != nil {
		return ctx.JSON(http.StatusNotFound, response.ErrorResponse{StatusCode: http.StatusNotFound, Message: "Database error"})
	}
	return ctx.JSON(http.StatusNotFound, response.SuccessResponse{StatusCode: http.StatusOK, Message: "User deleted"})
}
