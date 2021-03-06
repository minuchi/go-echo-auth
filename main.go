package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minuchi/go-echo-auth/controllers"
	"github.com/minuchi/go-echo-auth/database"
	"github.com/minuchi/go-echo-auth/lib"
	"github.com/minuchi/go-echo-auth/lib/middlewares"
	"github.com/minuchi/go-echo-auth/models"
	"net/http"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	config := lib.LoadConfig()

	jwtConfig := middleware.JWTConfig{
		SigningKey:     []byte(config.Keys.AccessTokenSecret),
		SuccessHandler: middlewares.JWTSuccessHandler,
	}

	db := database.Connect(config.Database)
	models.Migrate(db)

	e.Use(middlewares.ContextDB(db))

	g := e.Group("/api/auth/v1")
	g.GET("/time", controllers.GetTime)
	g.POST("/login", controllers.Login)
	g.POST("/signup", controllers.SignUp)
	g.POST("/token", controllers.IssueAccessToken)

	// Authorized routes
	g.Use(middleware.JWTWithConfig(jwtConfig))
	g.POST("/verify", controllers.Verify)
	e.Logger.Fatal(e.Start(":8080"))
}
