package main

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris"
)

var db *gorm.DB

func main() {

	var err error

	db, err = gorm.Open("mysql", "root:borodin@/test")

	defer db.Close()

	if err != nil {
		log.Print(err)
	}

	app := iris.New()

	api := app.Party("/api")
	{
		api.Get("/posts", getHandler)
		api.Post("/posts", addHandler)
	}

	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)

	app.Run(iris.Addr(":8080"))
}

