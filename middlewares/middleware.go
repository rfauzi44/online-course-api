package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rfauzi44/online-course-api/libs"
)

func Role(roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, libs.ResError("Authorization header missing"))
			}
			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

			validatedToken, err := libs.CheckToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, libs.ResError("Token is not valid"))
			}

			claims := validatedToken.Claims.(jwt.MapClaims)

			role := claims["role"].(string)
			allowed := false
			for _, r := range roles {
				if r == role {
					allowed = true
					break
				}
			}
			if !allowed {
				return c.JSON(http.StatusUnauthorized, libs.ResError("You do not have permission to access this resource"))
			}

			userID := claims["id"].(string)
			c.Set("authID", userID)
			return next(c)
		}
	}
}
