package library

import "fmt"

type Student struct{
	Name string
	Age int
}

func SayHello(name string){
	fmt.Println("hello")
	introduce(name)
}

func introduce(name string){
	fmt.Println("Nama Saya", name)
}