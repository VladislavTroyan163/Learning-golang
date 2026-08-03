package main

import (
	"fmt"
)

func main() {
	var c64 complex64 = 1.0 + 2.3i
	fmt.Println(c64)
	var c128 complex64 = 1.0 + 2.2i
	fmt.Println(c128)

	a := complex(1, 4)
	b := complex(2.4, 4.6)
	c := a * b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(real(c))
	fmt.Println(imag(c))

}
