package main

import (
	"../library"
	f "fmt"
)

func main(){
	var s1 = library.Student{"Fajar", 23}

	f.Println("Nama :",s1.Name)
	f.Println("Age :",s1.Age)
}