package dao

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return db
}

type Food struct {
	Id    int
	Title string
	Price float32
	Stock int
	Type  int
}

// 为Food绑定表名
func (v Food) TableName() string {
	return "foods"
}

// 增
func Add(food *Food) {
	db := GetDb()
	db.Create(food)
}

// 删
func Del(id int) {
	db := GetDb()
	db.Delete(Food{}, "id=?", id)
}

// 改
func UpdateData(food *Food) {

	db := GetDb()
	db.Exec("UPDATE foods SET price=?,stock=?,title=?,type=? WHERE id=?",
		food.Price, food.Stock, food.Title, food.Type, food.Id)

}

// 查所有
func Query() []Food {
	db := GetDb()
	//var pList []T_product
	var pList []Food
	db.Find(&pList)
	return pList
}

// 按id查
func Select(id int) Food {
	db := GetDb()
	var t_p Food
	db.Raw("SELECT * FROM foods WHERE id = ?", id).Scan(&t_p)
	return t_p
}
