package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/minuchi/go-echo-auth/lib"
	userService "github.com/minuchi/go-echo-auth/services/user"
	"gorm.io/gorm"
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
		PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
	}
	getAccessTokenRequest struct {
		RefreshToken string `json:"refresh_token"`
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

	db := c.Get("db").(*gorm.DB)
	email := body.Email
	userHashedPassword, err := userService.GetUserPasswordByEmail(db, email)
	if err != nil {
		return err
	}

	password := body.Password
	result, _ := lib.VerifyPassword(userHashedPassword, password)
	if result == false {
		return echo.ErrBadRequest
	}

	userId := userService.GetUserIdByEmail(db, email)
	refreshToken := lib.CreateRefreshToken(userId)

	return c.JSON(http.StatusOK, echo.Map{"refresh_token": refreshToken})
}

func SignUp(c echo.Context) error {
	body := new(signUpRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(body); err != nil {
		return err
	}

	db := c.Get("db").(*gorm.DB)
	email := body.Email

	userCount := userService.CheckUserExists(db, email)
	if userCount > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "This email already exists.")
	}
	hashedPassword := lib.HashPassword(body.Password)

	userService.CreateUser(db, email, hashedPassword)

	return c.JSON(http.StatusOK, body)
}

func IssueAccessToken(c echo.Context) error {
	body := new(getAccessTokenRequest)
	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(body); err != nil {
		return err
	}

	id := lib.DecryptRefreshToken(body.RefreshToken)

	accessToken := lib.CreateAccessToken(id)

	return c.JSON(http.StatusOK, echo.Map{"access_token": accessToken})
}

func Verify(c echo.Context) error {
	userId := c.Get("userId").(uint)
	return c.JSON(http.StatusOK, echo.Map{"ok": true, "user_id": userId})
}
