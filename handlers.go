package main

import "github.com/kataras/iris"

func notFoundHandler(ctx iris.Context) {
	ctx.HTML("<h1>Not found, sorry</h1>")
	ctx.StatusCode(404)
}

func getHandler(ctx iris.Context) {

	from := ctx.URLParam("from")
	count := ctx.URLParam("count")

	if from == "" {
		from = "0"
	}

	if count == "" {
		count = "20"
	}

	var post []Photo
	db.Where("id >= ?", from).Limit(count).Find(&post)

	ctx.JSON(post)
}

func addHandler(ctx iris.Context) {

	var post Photo

	ctx.ReadJSON(&post)
	db.Create(&post)
	db.Save(&post)

	ctx.JSON(post)
}
