package dao

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("error")
	}
	return db
}

type Food struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
	Type  int     `json:"type"`
}

// 为Food绑定表名
func (v Food) TableName() string {
	return "foods"
}

// 增
func Add(food *Food) {
	db := GetDb()
	if err := db.Create(food).Error; err != nil {
		// 错误处理
		fmt.Println("add error")
	}

}

// 删
func Del(id int) {
	db := GetDb()

	if err := db.Delete(Food{}, "id=?", id).Error; err != nil {
		// 错误处理
		fmt.Println("del error")
	}
}

// 改
func UpdateData(food *Food) {

	db := GetDb()

	if err := db.Exec("UPDATE foods SET price=?,stock=?,title=?,type=? WHERE id=?",
		food.Price, food.Stock, food.Title, food.Type, food.Id).Error; err != nil {
		// 错误处理
		fmt.Println("update error")
	}

}

// 查所有
func Query() []Food {
	db := GetDb()
	//var pList []T_product
	var pList []Food

	if err := db.Find(&pList).Error; err != nil {
		// 错误处理
		fmt.Println("selectAll error")
	}
	return pList
}

// 按id查
func Select(id int) Food {
	db := GetDb()
	var t_p Food

	if err := db.Raw("SELECT * FROM foods WHERE id = ?", id).Scan(&t_p).Error; err != nil {
		// 错误处理
		fmt.Println("select error")
	}
	return t_p
}
