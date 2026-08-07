package main

import "fmt"

func main() {
	x := 42
	// fmt.Println("Значение х:", x)
	// fmt.Println("Указатель на х:", &x)

	x2 := x
	fmt.Println("Значение х2:", x2)
	x2 = 100
	fmt.Println("Значение х2:", x2)
	fmt.Println("Значение х:", x)

	p := &x
	fmt.Println("Значение p:", p)
	fmt.Println("Значение  про аресу p:", *p)

	*p = 100
	fmt.Println("Значение  про аресу p:", *p)
	fmt.Println("Значение х:", x)
}
