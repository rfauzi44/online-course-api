package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/online-course-api/controllers"
	"github.com/rfauzi44/online-course-api/libs"
	"github.com/rfauzi44/online-course-api/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		data := map[string]string{
			"repository": "https://github.com/rfauzi44/online-course-api",
			"demo":       "https://online-course.zfdlzr.xyz",
			"postman":    "https://documenter.getpostman.com/view/25042327/2s93ebSVop",
		}
		return c.JSON(http.StatusOK, libs.ResSuccess("Hello there! This is online course API", data))
	})

	auth := e.Group("/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)

	public := e.Group("/public")
	public.GET("/course", controllers.ReadAllCourse)
	public.GET("/course/:id", controllers.GetCourseById)
	public.GET("/course/search", controllers.SearchCourse)
	public.GET("/course/sort", controllers.SortCourse)
	public.GET("/course/category", controllers.GetCategory)
	public.GET("/course/popular", controllers.GetPopularCategory)

	admin := e.Group("/admin", middlewares.Role("admin"))
	admin.POST("/course", controllers.CreateCourse)
	admin.PUT("/course/:id", controllers.UpdateCourse)
	admin.DELETE("/course/:id", controllers.DeleteCourse)
	admin.GET("/user", controllers.ReadAllUser)
	admin.DELETE("/user", controllers.DeleteUser)
	admin.GET("/stats", controllers.GetStats)
	//for testing //change role
	public.PUT("/for-testing", controllers.ChangeRole)

	return e
}
