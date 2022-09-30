package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
    ID           int64
    Name         string
    Age          int64
}
func (u User) TableName() string {
    //绑定MYSQL表名为users
    return "users"
}
func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
	user := User{Name: "q1mi", Age: 18}
 

db.Create(&user)   // 创建user
fmt.Println(db.Migrator().HasTable("users"))var user User



}
