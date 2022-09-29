package read

import (
	"fmt"
	"net/http"
	"project/dao"
	"project/service"
	"strconv"

	"github.com/labstack/echo"
)

func Add(c echo.Context) (err error) {
	u := new(dao.Food)
	if err = c.Bind(u); err != nil {

		return c.JSON(http.StatusOK, "add fail")
	}
	service.Save(u)
	return c.JSON(http.StatusOK, "add success")
}

// 删除
func Del(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusOK, "del fail")
	}
	i, _ := strconv.Atoi(id)

	service.Del(i)
	return c.JSON(http.StatusOK, "del success")
}

// 更新
func Update(c echo.Context) (err error) {

	product := new(dao.Food)
	if err = c.Bind(product); err != nil {
		return c.JSON(http.StatusOK, "update fail")
	}
	// 保存
	fmt.Println(product)
	service.Update(product)
	return c.JSON(http.StatusOK, "update success")
}

// 查询
func GetList(c echo.Context) error {
	param := c.QueryParam("id")

	if param == "" {
		return c.JSON(http.StatusOK, service.SelectAll())
	}
	i, _ := strconv.Atoi(param)
	return c.JSON(http.StatusOK, service.Select(i))
}
