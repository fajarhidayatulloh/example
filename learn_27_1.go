// +build main3

package main

func main() {
	/*var secret interface{}

	secret = "Kopi Gula"

	f.Println(secret)

	secret = []string{"Apel", "Mangga", "Melon"}
	f.Println(secret)

	secret = 12.4
	f.Println(secret)*/

	var data map[string]interface{}

	data = map[string]interface{}{
		"nama":      "Kopi Gula",
		"age":       23,
		"breakfast": []string{"Apel", "mangga", "melon"},
	}
}
