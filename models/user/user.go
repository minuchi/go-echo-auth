package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"index:,unique,where: deleted_at is null" json:"email" validate:"email"`
	Password string `json:"string"`
}
