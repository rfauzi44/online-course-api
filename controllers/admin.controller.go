package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/online-course-api/libs"
	"github.com/rfauzi44/online-course-api/models"
)

func GetStats(c echo.Context) error {
	data, err := models.GetStats()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	return c.JSON(http.StatusOK, libs.ResSuccess("success", data))

}
