package main

import (
	json "encoding/json"
	f "fmt"
	http "net/http"
)

type student struct {
	ID   string
	Nama string
	Age  int
}

var data = []student{
	student{"E001", "Fajar", 23},
	student{"E002", "IJKL", 24},
	student{"E003", "EFGH", 25},
	student{"E004", "ABCD", 26},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var id = r.FormValue("id")
		var result []byte
		var err error

		for _, each := range data {
			if each.ID == id {
				result, err = json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				w.Write(result)
				return
			}

			http.Error(w, "User tidak ada ", http.StatusBadRequest)
			return
		}

		http.Error(w, "", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	f.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
