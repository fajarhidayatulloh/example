package main

import "fmt"

func main() {
	var nilai = 8

	if nilai >= 8 {
		fmt.Println("Lulus")
	} else if nilai >= 5 && nilai <= 4 {
		fmt.Println("Lulus Bersyarat")
	} else {
		fmt.Println("Tidak Lulus")
	}
}
