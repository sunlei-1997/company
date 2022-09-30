package dao

import (
	"project/entity"
	"reflect"
	"testing"
)

var foodDao = new(FoodDao)

func TestSelect(t *testing.T) {
	// 程序输出的结果
	got := foodDao.Select(4)
	// 期望的结果
	want := &entity.Food{
		Id:    4,
		Title: "sunlei",
		Price: 1233,
		Stock: 23,
		Type:  31,
	}
	// 因为struct不能比较直接，借助反射包中的方法比较
	if !reflect.DeepEqual(want, got) {
		// 测试失败输出错误提示
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
func TestAdd(t *testing.T) {
	// 程序输出的结果
	data := &entity.Food{
		Title: "4",
		Price: 4,
		Stock: 4,
		Type:  4,
	}
	foodDao.Add(data)
	got := foodDao.Select(11)
	// 期望的结果
	want := &entity.Food{
		Id:    11,
		Title: "4",
		Price: 4.0,
		Stock: 4,
		Type:  4,
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
func TestUpdate(t *testing.T) {
	// 程序输出的结果
	data := &entity.Food{
		Id:    11,
		Title: "sunleiupdate",
		Price: 1233,
		Stock: 23,
		Type:  31,
	}
	foodDao.UpdateData(data)
	got := foodDao.Select(11)
	// 期望的结果
	want := &entity.Food{
		Id:    11,
		Title: "sunleiupdate",
		Price: 1233,
		Stock: 23,
		Type:  31,
	}
	// 因为struct不能比较直接，借助反射包中的方法比较
	if !reflect.DeepEqual(want, got) {
		// 测试失败输出错误提示
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
func TestDel(t *testing.T) {
	// 程序输出的结果
	foodDao.Del(2)
	got := foodDao.Select(2)
	// 期望的结果
	// 因为struct不能比较直接，借助反射包中的方法比较
	if got != nil {
		// 测试失败输出错误提示
		t.Errorf("expected:%v, got:%v", nil, got)
	}
}
