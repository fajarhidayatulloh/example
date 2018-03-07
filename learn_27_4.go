// +build main3

package main

import (
	f "fmt"
)

func main() {
	var person = []map[string]interface{}{
		{"nama": "Fajar", "age": 23},
		{"nama": "Gokil", "age": 24},
		{"nama": "Mantap", "age": 25},
	}

	for _, row := range person {
		f.Println("Umur", row["nama"], "Adalah", row["age"])
	}
}
