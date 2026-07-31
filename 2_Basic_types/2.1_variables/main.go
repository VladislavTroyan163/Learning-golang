package main

import (
	"fmt"
)

const (
	A = 1
	B
	C
)

func main() {
	// var age int = 30
	// var age = 30
	age := 30
	fmt.Println(age)

	// var x, y int = 1, 2
	x, y := 3, 4
	fmt.Println(x, y)
	y = 5
	fmt.Println(x, y)

	// const PI float64 = 3.14
	const PI = 3.14
	fmt.Println(PI)

	fmt.Println(A, B, C)
}
