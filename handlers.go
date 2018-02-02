package main

import "github.com/kataras/iris"

func notFoundHandler(ctx iris.Context) {
	ctx.HTML("<h1>Not found, sorry</h1>")
	ctx.StatusCode(404)
}

func getHandler(ctx iris.Context) {

	id := ctx.URLParam("id")

	if (id == "") {
		var country []Country
		db.Find(&country)
		ctx.JSON(country)
	} else {
		var country Country
		db.Where("id = ?", id).Find(&country)
		ctx.JSON(country)
	}

}

func addHandler(ctx iris.Context) {

	var action Action

	ctx.ReadJSON(&action)

	var country Country

	db.First(&country, action.ID) // find product with id 1

	newCount := country.Count

	if (action.Type == 1) {
		newCount = country.Count + 1
	} else {
		if (country.Count > 0) {
			newCount = country.Count - 1
		}
	}

	db.Model(&country).Update("count", newCount)

	ctx.JSON(country)
}
