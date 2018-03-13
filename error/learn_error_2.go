// +build main3

package main

import (
	err "errors"
	f "fmt"
	s "strings"
)

func validate(input string) (bool, error) {
	if s.TrimSpace(input) == "" {
		return false, err.New("Tidak Boleh Kosong")
	}
	return true, nil
}

func main() {
	var nama string
	f.Print("Masukan Nama Anda : ")
	f.Scanln(&nama)

	if valid, er := validate(nama); valid {
		f.Println("Hello", nama)
	} else {
		panic(er.Error())
	}
}
