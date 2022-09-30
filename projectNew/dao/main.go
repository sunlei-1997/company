package main

import (
	"net/http"
	//导入echo包
	"github.com/labstack/echo/v4"
	// "fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"projectNew/dao"
)

type User struct {
    ID           int64
    Name         string
    Age          int64
}

	


func main() {
	
	result := map[string]interface{}{}
	db := dao.GetDb() 
	db.Model(&User{}).First(&result)
	// fmt.Println(result)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		
		return c.JSON(http.StatusOK, result)
	})
	e.Logger.Fatal(e.Start(":8082"))
	
}

