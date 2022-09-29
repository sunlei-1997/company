package main

import (
	// "net/http"
	"project/controller/read"

	"github.com/labstack/echo"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	e := echo.New()
	read.ProductRouter(e)
	// e.Logger.Fatal(e.Start(":1323"))

	e.Logger.Fatal(e.Start(":1323"))
}
