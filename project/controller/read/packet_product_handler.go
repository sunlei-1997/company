package read

import (
	"net/http"
	"project/dao"
	"project/response"
	"project/service"
	"strconv"

	"github.com/labstack/echo"
)

func Add(c echo.Context) (err error) {
	u := new(dao.Food)
	if err = c.Bind(u); err != nil {
		return c.JSONBlob(http.StatusBadRequest, nil)
	}
	service.Save(u)
	return c.JSON(http.StatusOK, response.OK())
}

// 删除
func Del(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusOK, response.ParamError)
	}
	i, e := strconv.Atoi(id)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(nil))
	}
	service.Del(i)
	return c.JSON(http.StatusOK, response.OK())
}

// 更新
func Update(c echo.Context) (err error) {

	product := new(dao.Food)

	if err = c.Bind(product); err != nil {
		return c.JSONBlob(http.StatusBadRequest, nil)
	}
	// 保存
	service.Update(product)
	return c.JSON(http.StatusOK, response.OK())
}

// 查询
func GetList(c echo.Context) error {
	param := c.QueryParam("id")

	if param == "" {
		return c.JSON(http.StatusOK, response.All(service.SelectAll()))
	}
	i, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(nil))
	}
	return c.JSON(http.StatusOK, response.Single(service.Select(i)))
}
