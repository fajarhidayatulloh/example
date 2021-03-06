// +build main3

package main

import (
	sql "database/sql"
	f "fmt"

	_ "github.com/go-sql-driver/mysql"
)

/**
 * @Author: fajar
 * @Date: 2018-03-27 13:53:48
 * @Desc:
 */
type users struct {
	id       int
	username string
	name     string
	email    string
}

/**
 * @Author: fajar
 * @Date: 2018-03-27 13:54:16
 * @Desc:
 */
func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/lapnet")
	if err != nil {
		return nil, err
	}

	return db, nil
}

/**
 * @Author: fajar
 * @Date: 2018-03-27 13:54:22
 * @Desc:
 */
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

	rows, err := db.Query("select id, name, email from users")
	if err != nil {
		f.Println(err.Error())
		f.Println("ada error")
		return
	}
	defer rows.Close()

	var result []users

	for rows.Next() {
		var each = users{}
		var err = rows.Scan(&each.id, &each.name, &each.email)
		if err != nil {
			f.Println(err.Error())
			f.Println("error cuy")
			return
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		f.Println(err.Error())
		f.Println("error cuy")
		return
	}

	for _, each := range result {
		f.Println("id \t:", each.id)
		f.Println("Nama \t:", each.name)
		f.Println("email \t:", each.email)
		f.Println()
	}
}

/**
 * @Author: fajar
 * @Date: 2018-03-27 13:54:30
 * @Desc:
 */
func main() {
	SqlQuery()
}
