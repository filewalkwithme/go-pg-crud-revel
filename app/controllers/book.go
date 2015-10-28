package controllers

import "github.com/revel/revel"

//Book foo
type Book struct {
	*revel.Controller
}

//Index foo
func (b Book) Index() revel.Result {
	return b.Render()
}
