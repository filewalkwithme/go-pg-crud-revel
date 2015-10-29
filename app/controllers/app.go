package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	books, err := allBooks()
	if err != nil {
		panic(err)
	}

	return c.Render(books)
}
