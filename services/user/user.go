package user

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	userModel "github.com/minuchi/go-echo-auth/models/user"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, email, hashedPassword string) {
	user := userModel.User{Email: email, Password: hashedPassword}

	result := db.Create(&user)

	if err := result.Error; err == nil {
		fmt.Println(user)
	} else {
		fmt.Println(err)
	}
}

func CheckUserExists(db *gorm.DB, email string) int64 {
	var count int64
	user := userModel.User{}

	result := db.Model(&user).Where("email = ? AND deleted_at IS NULL", email).Count(&count)

	if err := result.Error; err == nil {
		return count
	} else {
		fmt.Println(err)
		return 0
	}
}

func GetUserIdByEmail(db *gorm.DB, email string) (id uint) {
	user := userModel.User{}
	result := db.Where("email = ? AND deleted_at IS NULL", email).First(&user)

	if err := result.Error; err == nil {
		id = user.ID
	}
	return
}

func GetUserPasswordByEmail(db *gorm.DB, email string) (string, error) {
	var user userModel.User
	result := db.Where("email = ?", email).First(&user)

	if err := result.Error; err == nil {
		return user.Password, nil
	} else if err == gorm.ErrRecordNotFound {
		return "", echo.ErrBadRequest
	} else {
		log.Error(err)
		return "", echo.ErrInternalServerError
	}
}
