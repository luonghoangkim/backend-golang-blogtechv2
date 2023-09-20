package middleware

import (
	"backend-blogtechv2/model"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		if user == nil {
			// Handle the error appropriately here, e.g. return an error response
			return c.JSON(http.StatusUnauthorized, model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "No user token provided",
				Data:       nil,
			})
		}

		userToken, ok := user.(*jwt.Token)
		if !ok {
			// Handle the error appropriately here, e.g. return an error response
			return c.JSON(http.StatusUnauthorized, model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid user token",
				Data:       nil,
			})
		}

		claims := userToken.Claims.(*model.JwtCustomClaim)
		if claims.Role != "ADMIN" {
			return c.JSON(http.StatusForbidden, model.Response{
				StatusCode: http.StatusForbidden,
				Message:    "Chỉ có quyền ADMIN mới được phép thực hiện tác vụ này",
				Data:       nil,
			})
		}

		return next(c)
	}
}
