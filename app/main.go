package main

import (
	// "net/http"
	"database/sql"
	"fmt"
	"guruchan-back/app/routes"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"

	_ "github.com/go-sql-driver/mysql"
)

// var baseurl string

func main() {
	//初始化
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.SetOutput(logging.LoggingSetting(os.Getenv("LogFile"))
	// echo实例
	e := echo.New() //new需要大写!!
	routes.RegisterRoutes(e)
	//start
	e.Start(":8080")

	//拿环境变量
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dsn := os.Getenv("DSN")
	if os.Getenv("ENV") == "production" {
		dsn = os.Getenv("PRODUCTION_DSN")
	}

	//db初始化
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("db连接失败")
		log.Fatal(err)
	}
	defer db.Close()

}
