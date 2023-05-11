package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/online-course-api/libs"
	"github.com/rfauzi44/online-course-api/models"
)

func CreateCourse(c echo.Context) error {
	var course models.Course
	err := c.Bind(&course)
	if err != nil {
		return err
	}
	v := validator.New()
	err = v.Struct(course)
	if err != nil {
		return c.JSON(http.StatusBadRequest, libs.ResError(err.Error()))
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}
	defer src.Close()

	url, image_id, err := libs.UploadImage(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	course.Image = url
	course.ImageID = image_id

	course.AuthorID = c.Get("authID").(string)

	data, err := models.CreateCourse(course)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}
	return c.JSON(http.StatusCreated, libs.ResSuccess("success", data))
}

func ReadAllCourse(c echo.Context) error {
	data, err := models.ReadAllCourse()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	return c.JSON(http.StatusOK, libs.ResSuccess("success", data))

}

func UpdateCourse(c echo.Context) error {

	var course models.Course
	course.ID = c.Param("id")
	err := c.Bind(&course)
	if err != nil {
		return err
	}

	v := validator.New()
	err = v.Struct(course)
	if err != nil {
		return c.JSON(http.StatusBadRequest, libs.ResError(err.Error()))
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}
	defer src.Close()

	url, image_id, err := libs.UploadImage(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	course.Image = url
	course.ImageID = image_id
	data, err := models.UpdateCourse(course)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}
	return c.JSON(http.StatusOK, libs.ResSuccess("success", data))
}

func DeleteCourse(c echo.Context) error {
	id := c.Param("id")
	data, err := models.DeleteCourse(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	return c.JSON(http.StatusOK, libs.ResSuccess(fmt.Sprintf("course with id %s deleted", id), data))
}

func GetCourseById(c echo.Context) error {
	id := c.Param("id")
	data, err := models.GetCourseById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	return c.JSON(http.StatusOK, libs.ResSuccess("successs", data))

}

func SearchCourse(c echo.Context) error {
	query := c.QueryParam("keyword")
	data, err := models.SearchCourse(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	return c.JSON(http.StatusOK, libs.ResSuccess("success", data))

}

func SortCourse(c echo.Context) error {
	query := c.QueryParam("price")
	data, err := models.SortCourse(query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	return c.JSON(http.StatusOK, libs.ResSuccess("success", data))

}

func GetCategory(c echo.Context) error {
	data, err := models.GetCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	return c.JSON(http.StatusOK, libs.ResSuccess("success", data))

}

func GetPopularCategory(c echo.Context) error {
	data, err := models.GetPopularCategory()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, libs.ResError(err.Error()))
	}

	return c.JSON(http.StatusOK, libs.ResSuccess("success", data))

}
