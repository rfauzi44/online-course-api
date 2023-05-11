package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/online-course-api/libs"
	"github.com/rfauzi44/online-course-api/models"
)

func ReadAllUser(c echo.Context) error {
	data, err := models.ReadAllUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	return c.JSON(http.StatusOK, libs.ResSuccess("success", data))

}

type UserDelete struct {
	Email string `json:"email" validate:"required,email"`
}

func DeleteUser(c echo.Context) error {
	var user UserDelete
	if err := c.Bind(&user); err != nil {
		return err
	}

	v := validator.New()
	err := v.Struct(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, libs.ResError(err.Error()))
	}
	data, err := models.DeleteUser(user.Email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}
	return c.JSON(http.StatusOK, libs.ResSuccess(fmt.Sprintf("user with email %s deleted", user.Email), data))
}

func ChangeRole(c echo.Context) error {

	email := c.FormValue("email")
	v := validator.New()
	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, libs.ResError(err.Error()))
	}
	data, err := models.ChangeRole(email)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}
	return c.JSON(http.StatusOK, libs.ResSuccess(fmt.Sprintf("user with email %s become admin", email), data))
}
