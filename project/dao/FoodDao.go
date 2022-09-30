package dao

import (
	"fmt"
	"project/cgo"
	"project/entity"
)

type FoodDao struct {
}

// 增
func (f *FoodDao) Add(food *entity.Food) {
	db := cgo.GetDb()
	if err := db.Create(food).Error; err != nil {
		// 错误处理
		fmt.Println("add error")
	}

}

// 删
func (f *FoodDao) Del(id int) {
	db := cgo.GetDb()

	if err := db.Delete(entity.Food{}, "id=?", id).Error; err != nil {
		// 错误处理
		fmt.Println("del error")
	}
}

// 改
func (f *FoodDao) UpdateData(food *entity.Food) {

	db := cgo.GetDb()

	if err := db.Exec("UPDATE foods SET price=?,stock=?,title=?,type=? WHERE id=?",
		food.Price, food.Stock, food.Title, food.Type, food.Id).Error; err != nil {
		// 错误处理
		fmt.Println("update error")
	}

}

// 查所有
func (f *FoodDao) Query() []*entity.Food {
	db := cgo.GetDb()
	//var pList []T_product
	var pList []*entity.Food

	if err := db.Find(&pList).Error; err != nil {
		// 错误处理
		fmt.Println("selectAll error")
	}
	return pList
}

// 按id查
func (f *FoodDao) Select(id int) *entity.Food {
	db := cgo.GetDb()
	var t_p *entity.Food

	if err := db.Raw("SELECT * FROM foods WHERE id = ?", id).Scan(&t_p).Error; err != nil {
		// 错误处理
		fmt.Println("select error")
	}
	return t_p
}
