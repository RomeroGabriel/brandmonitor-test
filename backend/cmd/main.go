package main

import (
	"github.com/RomeroGabriel/brandmonitor-challange/backend/internal/webserver/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler := handlers.NewSearchHandler()
	e.GET("/", handler.Search)

	e.Logger.Fatal(e.Start(":8080"))
}
