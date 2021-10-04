package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=db user=auth password=auth dbname=auth port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
