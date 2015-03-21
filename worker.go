package main

import (
	"fmt"
	"os"

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
	Id     int64  `db:"pk" column:"tbl_id"`
	UserId string `default:""`
	Title  string `default:""`
	Body   string `default:""`
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
	db := initDB()
	db.SetLogOutput(os.Stdout)
	var results []User
	sampleUsers := []User{
		{Name: "achiku", Active: true},
		{Name: "moqada", Active: false},
	}
	if _, err := db.Insert(sampleUsers); err != nil {
		panic(err)
	}
	if err := db.Select(&results); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", results)
	if _, err := db.Delete(&results); err != nil {
		panic(err)
	}
	if err := db.Select(&results, db.Where("active", "=", false)); err != nil {
		panic(err)
	}
	if _, err := db.Delete(&results); err != nil {
		panic(err)
	}
	fmt.Printf("done\n")
}
