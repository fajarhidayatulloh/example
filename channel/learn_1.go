package main

import (
	f "fmt"
	r "runtime"
)

func main() {
	r.GOMAXPROCS(2)

	var message = make(chan string)

	var sayHelloTo = func(who string) {
		var data = f.Sprintf("Hello %s", who)
		message <- data
	}

	go sayHelloTo("Fajar")
	go sayHelloTo("Hidayat")
	go sayHelloTo("Gokil")

	var message1 = <-message
	f.Println(message1)

	var message2 = <-message
	f.Println(message2)

	var message3 = <-message
	f.Println(message3)
}
