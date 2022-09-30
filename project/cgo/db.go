package cgo

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("error")
	}

	return db
}
