// +build main3

package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var allstudent = []person{
		{name: "Fajar", age: 20},
		{name: "Hidayat", age: 21},
		{name: "Gokil", age: 22},
	}

	for _, student := range allstudent {
		fmt.Println(student.name, "age is ", student.age)
	}
}
