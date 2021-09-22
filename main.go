package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minuchi/go-echo-auth/controllers"
)

func main() {
	e := echo.New()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	g := e.Group("/api/auth/v1")
	g.GET("/time", controllers.GetTime)
	g.POST("/login", controllers.Login)
	e.Logger.Fatal(e.Start(":8080"))
}
