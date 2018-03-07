// +build main3

package main

import (
	f "fmt"
	r "reflect"
)

type person struct {
	Nama string
	Age  int
}

func (s *person) getPropertyInfo() {
	var reflectValue = r.ValueOf(s)

	if reflectValue.Kind() == r.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		f.Println("nama \t\t:", reflectType.Field(i).Name)
		f.Println("tipe data \t:", reflectType.Field(i).Type)
		f.Println("nilai \t\t:", reflectValue.Field(i).Interface())
		f.Println("")
	}
}

func main() {
	var s1 = &person{Nama: "Fajar", Age: 23}
	s1.getPropertyInfo()
}
