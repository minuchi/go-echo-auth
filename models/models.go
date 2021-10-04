package models

import (
	"github.com/minuchi/go-echo-auth/models/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&user.User{})
}
