package router

import (
	"github.com/kataras/iris/v12"
	"iris-template/controller"
)

func books(app *iris.Application)  {
	api := app.Party("/books")

	// GET: http://localhost:8080/books
	api.Get("/", controller.BooksList)
	// POST: http://localhost:8080/books
	api.Post("/", controller.BooksCreate)
}

func Init(app *iris.Application) {
	books(app)
}
