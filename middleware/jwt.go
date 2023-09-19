package middleware

import (
	"backend-blogtechv2/model"
	"backend-blogtechv2/security"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaim{},
		SigningKey: []byte(security.Secret_key),
	}

	return middleware.JWTWithConfig(config)
}