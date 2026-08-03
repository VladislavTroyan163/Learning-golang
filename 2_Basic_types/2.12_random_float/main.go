package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// random := rand.Float64()
	// fmt.Println(random)

	min := 20.0
	max := 50.0
	random := rand.Float64()*(max-min) + min
	fmt.Println(random)
}
