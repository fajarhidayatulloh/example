// +build main3

package main

import (
	sql "database/sql"
	json "encoding/json"
	f "fmt"
	http "net/http"

	_ "github.com/go-sql-driver/mysql"
)

type users struct {
	Id       int
	Name     string
	Username string
	//email    string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/lapnet")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func showJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := connect()
	if err != nil {
		f.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, name, email from users")
	if err != nil {
		f.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []users

	for rows.Next() {
		var each = users{}
		var err = rows.Scan(&each.Id, &each.Name, &each.Username)
		if err != nil {
			f.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		f.Println(err.Error())
		return
	}

	if r.Method == "GET" {

		hasil, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(hasil)
		return

	}

	http.Error(w, "", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/users", showJson)

	f.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
