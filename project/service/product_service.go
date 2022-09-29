package service

import (
	"project/dao"
)

func Save(do *dao.Food) {
	dao.Add(do)
}
func Del(id int) {
	dao.Del(id)
}
func SelectAll() []dao.Food {
	return dao.Query()
}
func Select(id int) dao.Food {
	return dao.Select(id)
}
func Update(do *dao.Food) {
	dao.UpdateData(do)
}
