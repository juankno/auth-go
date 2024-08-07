package controllers

import (
	"go-auth/models"
	"go-auth/services"
	"go-auth/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService}
}

func (c *AuthController) Register(ctx echo.Context) error {
	user := new(models.User)
	if err := ctx.Bind(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := utils.ValidateStruct(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, utils.FormatValidationError(err))
	}

	if err := c.authService.Register(user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to register user"})
	}

	return ctx.JSON(http.StatusCreated, user)
}

func (c *AuthController) Login(ctx echo.Context) error {
	var login struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	if err := ctx.Bind(&login); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if err := utils.ValidateStruct(&login); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	token, err := c.authService.Login(login.Email, login.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	return ctx.JSON(http.StatusOK, map[string]string{"token": token})
}
