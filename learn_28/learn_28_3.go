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

func (s *person) SetName(nama string) {
	s.Nama = nama
}

func main() {
	var s1 = &person{Nama: "Fajar Hidayat", Age: 23}
	f.Println("Nama \t:", s1.Nama)

	var reflectValue = r.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]r.Value{
		r.ValueOf("Hidayat"),
	})
	f.Println("Nama \t:", s1.Nama)
}
