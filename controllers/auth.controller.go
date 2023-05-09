package controllers

import (
	"net/http"

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
	hashPassword, err := libs.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	data, err := models.Register(user.Email, hashPassword)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	return c.JSON(http.StatusOK, libs.ResSuccess("Registered", data))

}

func Login(c echo.Context) error {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	data, err := models.Login(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, libs.ResError(err.Error()))
	}

	token, err := libs.GenerateToken(data.ID, data.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, libs.ResSuccess("You're login", token))

}
