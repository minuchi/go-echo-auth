package lib

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTTokenClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}

var (
	refreshTokenSecret string
	accessTokenSecret  string
)

func init() {
	config := LoadConfig()

	refreshTokenSecret = config.Keys.RefreshTokenSecret
	accessTokenSecret = config.Keys.AccessTokenSecret
}

func createToken(secret string, claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	if err != nil {
		panic(err)
	}

	return t, err
}

func CreateRefreshToken(id uint) string {
	claims := &JWTTokenClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
		},
	}

	t, _ := createToken(refreshTokenSecret, claims)

	return t
}

func CreateAccessToken(id uint) string {
	claims := &JWTTokenClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}

	t, _ := createToken(accessTokenSecret, claims)

	return t
}

func decryptJWTToken(secret, tokenString string) uint {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	id := uint(claims["id"].(float64))

	return id
}

func DecryptRefreshToken(refreshToken string) (id uint) {
	id = decryptJWTToken(refreshTokenSecret, refreshToken)
	return
}
