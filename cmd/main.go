package main

import (
	"go-auth/config"
	"go-auth/routes"
	"go-auth/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	db := config.InitDB()
	defer db.Close()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.InitRoutes(e, db)

	port := utils.GetPort()
	e.Logger.Fatal(e.Start(port))
}
