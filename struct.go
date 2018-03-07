// +build main3

package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	grade int
	person
}

func main() {
	var p1 = person{name: "Fajar", age: 23}

	var s1 = student{person: p1, grade: 3}

	fmt.Println("Nama : ", s1.name)
	fmt.Println("Grade : ", s1.grade)
	fmt.Println("Age : ", s1.age)
}
