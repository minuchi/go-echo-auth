package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type (
	loginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
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

	return c.JSON(http.StatusOK, body)
}
