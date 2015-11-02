package controllers

import (
	"time"
	"strconv"

	"github.com/revel/revel"
	"github.com/maiconio/go-pg-crud-revel/app/models"
	"github.com/maiconio/go-pg-crud-revel/app/routes"
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

func (c App) SaveBook() revel.Result {
	var err error
	var id = 0

	idStr := c.Params.Get("id")

	if len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}
	}

	name := c.Params.Get("name")
	author := c.Params.Get("author")

	pagesStr := c.Params.Get("pages")
	pages := 0
	if len(pagesStr) > 0 {
		pages, err = strconv.Atoi(pagesStr)
		if err != nil {
			panic(err)
		}
	}

	publicationDateStr := c.Params.Get("publicationDate")
	var publicationDate time.Time

	if len(publicationDateStr) > 0 {
		publicationDate, err = time.Parse("2006-01-02", publicationDateStr)
		if err != nil {
			panic(err)
		}
	}

	if id == 0 {
		_, err = insertBook(name, author, pages, publicationDate)
	} else {
		_, err = updateBook(id, name, author, pages, publicationDate)
	}

	if err != nil {
		panic(err)
	}

	return c.Redirect(routes.App.Index())
}

func (c App) DeleteBook() revel.Result {
	idStr := c.Params.Get("id")

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			panic(err)
		}

		_, err = removeBook(id)
		if err != nil {
			panic(err)
		}
	}

	return c.Redirect(routes.App.Index())
}
