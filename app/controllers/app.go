package controllers

import (
	"time"
	"strconv"

	"github.com/revel/revel"
	"github.com/maiconio/go-pg-crud-revel/app/models"
)

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

func (c App) Book() revel.Result {
	var book = models.Book{}	

	idStr := c.Params.Get("id")
	book.PublicationDate = time.Now()
	
	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
		
		book, err = getBook(id)
		if err != nil {
			panic(err)
		}
	}

	return c.Render(book)
}
