package main

import (
	// "net/http"
	"project/controller"

	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	e := echo.New()
	controller.ProductRouter(e)
	// e.Logger.Fatal(e.Start(":1323"))

	e.Logger.Fatal(e.Start(":1323"))
}
