package main

import (
	"go-auth/config"
	"go-auth/routes"
	"go-auth/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := config.InitDB()
	defer db.Close()

	routes.InitRoutes(e, db)

	port := utils.GetPort()
	e.Logger.Fatal(e.Start(port))
}
