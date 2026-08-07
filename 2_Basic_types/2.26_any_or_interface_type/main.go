package main

import "fmt"

func main() {
	var val interface{} = "Привет друг!"
	//var val any = "Привет друг!"
	fmt.Println(val)
	val = 50
	fmt.Println(val)

	if str, ok := val.(string); ok {
		fmt.Println("Это строка:", str)
	} else {
		fmt.Println("Это не строка")
	}

	if i, ok := val.(int); ok {
		fmt.Println("Это int:", i)
	} else {
		fmt.Println("Это не int")
	}
}
