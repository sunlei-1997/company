package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	
)




func GetDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
