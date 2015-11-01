package models

import "time"

//Book represents a book object
type Book struct {
	ID              int
	Name            string
	Author          string
	PublicationDate time.Time
	Pages           int
}

//PublicationDateStr returns a sanitized Publication Date in the format YYYY-MM-DD
func (b Book) PublicationDateStr() string {
	return b.PublicationDate.Format("2006-01-02")
}
