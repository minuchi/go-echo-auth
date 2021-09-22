package controllers

import (
	"fmt"
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
	errResponse struct {
		Ok      bool   `json:"ok"`
		Error   string `json:"error"`
		Message string `json:"message"`
	}
)

func getRequiredErrResponse(field string) errResponse {
	err := fmt.Sprintf("%s_required", field)
	msg := fmt.Sprintf("%s required.", field)
	return errResponse{Ok: false, Error: err, Message: msg}
}

func GetTime(c echo.Context) error {
	t := &getTimeResponse{
		Time: time.Now().Format(time.RFC3339),
	}

	return c.JSON(http.StatusOK, t)
}

func Login(c echo.Context) error {
	body := make(map[string]interface{})
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	email := body["email"]
	if email == nil {
		return echo.NewHTTPError(http.StatusBadRequest, getRequiredErrResponse("email"))
	}

	password := body["password"]
	if password == nil {
		return echo.NewHTTPError(http.StatusBadRequest, getRequiredErrResponse("password"))
	}

	return c.JSON(http.StatusOK, body)
}
