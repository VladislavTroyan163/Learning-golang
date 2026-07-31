package main

import "fmt"

const (
	RED = iota
	GREEN
	BLUE
	ORANGE
)
const (
	RED1 = iota
	GREEN1
	BLUE1   = 8
	ORANGE1 = iota
	YELLOY1
)

const (
	RED2 = iota + 5
	GREEN2
	_
	ORANGE2
)

const (
	RED3 = iota * 2
	GREEN3
	BLUE3
	ORANGE3
)

func main() {
	fmt.Println(RED, GREEN, BLUE, ORANGE)
	fmt.Println("--------")
	fmt.Println(RED1, GREEN1, BLUE1, ORANGE1, YELLOY1)
	fmt.Println("--------")
	fmt.Println(RED2, GREEN2, ORANGE2)
	fmt.Println("--------")
	fmt.Println(RED3, GREEN3, BLUE3, ORANGE3)
}
