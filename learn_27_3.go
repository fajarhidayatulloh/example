// +build main3

package main

import (
	f "fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	var secret interface{} = &person{name: "Fajar", age: 23}
	var name = secret.(*person).name
	f.Println("nama saya \t :", name)
}
