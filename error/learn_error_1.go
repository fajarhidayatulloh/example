package main

import (
	f "fmt"
	str "strconv"
)

func main() {
	var input string
	f.Print("Masukan Nomor : ")
	f.Scanln(&input)

	var number int
	var err error
	number, err = str.Atoi(input)

	if err == nil {
		f.Println(number, "benar")

	} else {
		f.Println(input, "Bukan Nomor")
		f.Println(err.Error())
	}
}
