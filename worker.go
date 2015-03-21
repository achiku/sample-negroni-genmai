package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/naoina/genmai"
)

type Knot struct {
	URL string
	DB  *genmai.DB
}

type User struct {
	Id     int64  `db:"pk" column:"tbl_id"`
	Name   string `default:""`
	Active bool   `db:"-"`
}

type Note struct {
	Id    int64  `db:"pk" column:"tbl_id"`
	Title string `default:""`
	Body  string `default:""`
}

func initDB() *genmai.DB {
	db, err := genmai.New(&genmai.SQLite3Dialect{}, "./sample.db")
	if err != nil {
		panic(err)
	}
	if err := db.CreateTableIfNotExists(&User{}); err != nil {
		panic(err)
	}
	if err := db.CreateTableIfNotExists(&Note{}); err != nil {
		panic(err)
	}
	return db
}

func main() {
	word := "world"
	fmt.Println("hello, ", word)
	db := initDB()
	var results []User
	sampleUsers := []User{
		{Name: "achiku", Active: true},
		{Name: "moqada", Active: true},
	}
	deleteTarget := []User{
		{Name: "achiku"},
		{Name: "moqada"},
	}
	if _, err := db.Insert(sampleUsers); err != nil {
		panic(err)
	}
	if err := db.Select(&results); err != nil {
		panic(err)
	}
	if _, err := db.Delete(deleteTarget); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", results)
}
