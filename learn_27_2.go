// +build main3

package main

import (
	f "fmt"
	s "strings"
)

func main() {
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10
	f.Println(secret, "Multiple by 10 is", number)

	secret = []string{"Kopi", "Gula", "Coklat"}
	var gruits = s.Join(secret.([]string), ", ")
	f.Println(gruits, "is my favorite fruits")
}
