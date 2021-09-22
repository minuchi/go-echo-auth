package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/minuchi/go-echo-auth/lib"
	"net/http"
	"time"
)

type (
	loginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	signUpRequest struct {
		Email           string `json:"email" validate:"required,email"`
		Password        string `json:"password" validate:"required"`
		PasswordConfirm string `json:"passwordConfirm" validate:"required,eqfield=Password"`
	}

	getTimeResponse struct {
		Time string `json:"time"`
	}
)

func GetTime(c echo.Context) error {
	t := &getTimeResponse{
		Time: time.Now().Format(time.RFC3339),
	}

	return c.JSON(http.StatusOK, t)
}

func Login(c echo.Context) error {
	body := new(loginRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(body); err != nil {
		return err
	}

	// Verify password
	password := body.Password
	hashedPassword := lib.HashPassword(password)
	result, _ := lib.VerifyPassword(hashedPassword, password)

	if result == true {
		fmt.Printf("Password verified with %s\n", hashedPassword)
	}

	return c.JSON(http.StatusOK, body)
}

func SignUp(c echo.Context) error {
	body := new(signUpRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(body); err != nil {
		return err
	}

	email := body.Email
	hashedPassword := lib.HashPassword(body.Password)
	fmt.Printf("email: %s\nhashed_password: %s\n", email, hashedPassword)

	return c.JSON(http.StatusOK, body)
}
