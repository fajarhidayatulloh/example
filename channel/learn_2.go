// +build main3

package main

import (
	f "fmt"
	r "runtime"
)

func printMessage(what chan string) {
	f.Println(<-what)
}

func main() {
	r.GOMAXPROCS(2)

	var message = make(chan string)

	for _, row := range []string{"Fajar", "Hidayat", "Gokil"} {
		go func(who string) {
			var data = f.Sprintf("Hello %s", who)
			message <- data
		}(row)
	}

	for i := 0; i < 3; i++ {
		printMessage(message)
	}
}
