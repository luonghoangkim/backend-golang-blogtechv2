package middleware

import (
	"backend-blogtechv2/model"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenData, ok := c.Get("user").(*jwt.Token)
		if !ok || tokenData == nil {
			return c.JSON(http.StatusUnauthorized, model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Token không hợp lệ hoặc không tồn tại",
				Data:       nil,
			})
		}

		clams := tokenData.Claims.(*model.JwtCustomClaim)

		// Kiểm tra xem người dùng có quyền ADMIN không
		if clams.Role != "ADMIN" {
			return c.JSON(http.StatusUnauthorized, model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Chỉ có quyền ADMIN mới được phép thực hiện tác vụ này",
				Data:       nil,
			})
		}

		return next(c)
	}
}

