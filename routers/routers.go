package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/online-course-api/controllers"
	"github.com/rfauzi44/online-course-api/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		link := "<a href='https://documenter.getpostman.com/view/25042327/2s93eSYvCx'>here</a>"
		response := "Hello World! This is Online Course API. You can check Postman Documentation " + link
		return c.HTML(http.StatusOK, response)
	})

	e.GET("/testauth", func(c echo.Context) error {
		link := "<a href='https://documenter.getpostman.com/view/25042327/2s93eSYvCx'>here</a>"
		response := "Hello World! This is Online Course API. You can check Postman Documentation " + link
		return c.HTML(http.StatusOK, response)
	}, middlewares.AuthMiddleware("admin"))

	//Auth
	e.POST("/auth/register", controllers.Register)
	e.POST("/auth/login", controllers.Login)

	return e
}
