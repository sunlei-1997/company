package main

import (
	"fmt"
)

type Employee struct {
	id   string
	name string
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func main() {
	emp1 := &Employee{
		id:   "E1",
		name: "Shashank",
	}
	emp1.SetName("Vivek")
	fmt.Println(emp1.name)
}
