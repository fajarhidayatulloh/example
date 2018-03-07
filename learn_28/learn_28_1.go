package main

import (
	f "fmt"
	r "reflect"
)

func main() {
	var number = 23

	var reflectValue = r.ValueOf(number)

	f.Println("tipe variabel \t:", reflectValue.Type())
	f.Println("nilai variabel \t:", reflectValue.Interface())
	/*if reflectValue.Kind() == r.Int {
		f.Println("nilai variabel \t:", reflectValue.Int())
	}*/
}
