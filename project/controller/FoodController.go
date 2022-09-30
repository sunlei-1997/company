package controller

import (
	"net/http"
	"project/entity"
	"project/response"
	"project/service"
	"strconv"

	"github.com/labstack/echo"
)

type FoodController struct {
}

var productService = new(service.ProductService)

func (f *FoodController) Add(c echo.Context) (err error) {
	u := new(entity.Food)
	if err = c.Bind(u); err != nil {
		return c.JSONBlob(http.StatusBadRequest, nil)
	}
	productService.Save(u)
	return c.JSON(http.StatusOK, response.OK())
}

// 删除
func (f *FoodController) Del(c echo.Context) error {
	id := c.QueryParam("id")
	if id == "" {
		return c.JSON(http.StatusOK, response.ParamError)
	}
	i, e := strconv.Atoi(id)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(nil))
	}
	productService.Del(i)
	return c.JSON(http.StatusOK, response.OK())
}

// 更新
func (f *FoodController) Update(c echo.Context) (err error) {

	product := new(entity.Food)

	if err = c.Bind(product); err != nil {
		return c.JSONBlob(http.StatusBadRequest, nil)
	}
	// 保存
	productService.Update(product)
	return c.JSON(http.StatusOK, response.OK())
}

// 查询
func (f *FoodController) GetList(c echo.Context) error {
	param := c.QueryParam("id")

	if param == "" {
		return c.JSON(http.StatusOK, response.All(productService.SelectAll()))
	}
	i, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error(nil))
	}
	return c.JSON(http.StatusOK, response.Single(productService.Select(i)))
}
