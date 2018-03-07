package main

import (
	f "fmt"
	r "runtime"
)

func testGoRoutine(till int, message string) {
	for i := 0; i < till; i++ {
		f.Println((i + 1), message)
	}
}

func main() {
	r.GOMAXPROCS(2)

	go testGoRoutine(5, "Hello")
	testGoRoutine(5, "Apa Kabar")

	var input string
	f.Scanln(&input)
}
