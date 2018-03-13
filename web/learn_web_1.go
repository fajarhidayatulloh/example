package main

import (
	f "fmt"
	temp "html/template"
	web "net/http"
)

func index(w web.ResponseWriter, r *web.Request) {
	f.Fprintln(w, "Hello, Apa Kabar?")
}

func main() {
	web.HandleFunc("/", func(w web.ResponseWriter, r *web.Request) {
		//f.Fprintln(w, "Hallo")

		var data = map[string]string{
			"Nama": "Fajar",
			"Hobi": "Sepedaan",
		}

		var t, er = temp.ParseFiles("template.html")
		if er != nil {
			panic(er.Error())
		}

		t.Execute(w, data)
	})

	web.HandleFunc("/index", index)

	f.Println("Starting web server at http://localhost:8080/")
	web.ListenAndServe(":8080", nil)
}
