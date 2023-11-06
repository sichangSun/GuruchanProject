package main

import (
	// "net/http"
	"guruchan-back/app/routes"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

// var baseurl string

func main() {
	// echo实例
	e := echo.New() //new需要大写!!
	routes.RegisterRoutes(e)
	//start
	e.Start(":8080")
}
