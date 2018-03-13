// +build main3

package main

import (
	sql "database/sql"
	f "fmt"

	_ "github.com/go-sql-driver/mysql"
)

type users struct {
	id       int
	username string
	name     string
	email    string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/lapnet")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SqlQuery() {
	db, err := connect()
	if err != nil {
		f.Println(err.Error())
		f.Println("DB Gak Konek cuy")
		return
	} else {
		f.Println("DB Konek Cuy")
	}
	defer db.Close()
}

func main() {
	SqlQuery()
}
