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

func catch() {
	if r := recover(); r != nil {
		f.Println("error occurd", r)
	} else {
		f.Println("Applikasi berjalan dengan baik")
	}
}

func main() {
	defer catch()

	var nama string
	f.Print("Masukan Nama Anda : ")
	f.Scanln(&nama)

	if valid, er := validate(nama); valid {
		f.Println("Hallo", nama)
	} else {
		panic(er.Error())
		f.Println("selesai")
	}
}
