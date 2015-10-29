package controllers

import (
	"database/sql"
	"fmt"

	"github.com/revel/revel"
)

//DB represents the database instance
var db *sql.DB

//InitDB initializes DB connection
func InitDB() {
	connstring := fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable", "postgres", "", "books_database")

	var err error
	db, err = sql.Open("postgres", connstring)
	if err != nil {
		revel.INFO.Println("DB Error", err)
	}
	revel.INFO.Println("DB Connected")
}

func init() {
	revel.OnAppStart(InitDB)
}
