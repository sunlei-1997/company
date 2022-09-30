package controller

import (
	"github.com/labstack/echo"
)

var foodController = new(FoodController)

// 路由
func ProductRouter(e *echo.Echo) {
	// 增
	e.POST("/add", foodController.Add)
	// 删
	e.POST("/del", foodController.Del)
	// 改
	e.POST("/update", foodController.Update)
	// 查
	e.GET("/get", foodController.GetList)
}
