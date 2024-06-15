package middleware

import (
	jwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := jwt.Config{
		SigningKey: []byte("secret"),
	}

	return jwt.WithConfig(config)
}
