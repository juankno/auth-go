package routes

import (
	"go-auth/controllers"
	"go-auth/middleware"
	"go-auth/repository"
	"go-auth/services"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {

	userRepository := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)

	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)

	r := e.Group("/")
	r.Use(middleware.JWTMiddleware())
	r.GET("user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "You are authorized"})
	})
}
