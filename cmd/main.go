package main

import (
	"go-auth/config"
	"go-auth/controllers"
	"go-auth/middleware"
	"go-auth/repository"
	"go-auth/services"
	"go-auth/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := config.InitDB()
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)

	e.POST("/register", authController.Register)
	e.POST("/login", authController.Login)

	r := e.Group("/restricted")
	r.Use(middleware.JWTMiddleware())
	r.GET("/user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "You are authorized"})
	})

	port := utils.GetPort()

	e.Logger.Fatal(e.Start(port))
}
