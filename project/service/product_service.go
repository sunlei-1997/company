package service

import (
	"project/dao"
	"project/entity"
)

var foodDao = new(dao.FoodDao)

type ProductService struct {
}

func (p *ProductService) Save(do *entity.Food) {
	foodDao.Add(do)
}
func (p *ProductService) Del(id int) {
	foodDao.Del(id)
}
func (p *ProductService) SelectAll() []*entity.Food {
	return foodDao.Query()
}
func (p *ProductService) Select(id int) *entity.Food {
	food := foodDao.Select(id)
	return food
}
func (p *ProductService) Update(do *entity.Food) {
	foodDao.UpdateData(do)
}
