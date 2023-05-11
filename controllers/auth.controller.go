package controllers

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/online-course-api/libs"
	"github.com/rfauzi44/online-course-api/models"
)

func Register(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	v := validator.New()
	err = v.Struct(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, libs.ResError(err.Error()))
	}

	hashPassword, err := libs.HashingPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	user.Password = hashPassword

	data, err := models.Register(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	return c.JSON(http.StatusCreated, libs.ResSuccess("register success", data))

}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
type LoginRes struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

func Login(c echo.Context) error {
	var user LoginReq
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	v := validator.New()
	err = v.Struct(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, libs.ResError(err.Error()))
	}

	data, err := models.Login(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, libs.ResError(err.Error()))
	}

	token, err := libs.GenerateToken(data.ID, data.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	response := LoginRes{
		User:  *data,
		Token: token,
	}

	return c.JSON(http.StatusOK, libs.ResSuccess("login success", response))

}
