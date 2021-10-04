package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func JWTSuccessHandler(c echo.Context) {
	user := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userId := uint(user["id"].(float64))
	c.Set("userId", userId)
}
