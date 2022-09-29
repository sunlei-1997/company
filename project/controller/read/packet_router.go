package read

import (
	"github.com/labstack/echo"
)

// 路由
func ProductRouter(e *echo.Echo) {
	// 增
	e.POST("/add", Add)
	// 删
	e.POST("/del", Del)
	// 改
	e.POST("/update", Update)
	// 查
	e.GET("/get", GetList)
}
