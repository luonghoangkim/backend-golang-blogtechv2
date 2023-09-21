package middleware

import (
	"backend-blogtechv2/model"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AdminOnlyMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*model.JwtCustomClaim)
		role := claims.Role
		if role != "ADMIN" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Chỉ ADMIN mới được post")
		}
		return next(c)
	}
}
