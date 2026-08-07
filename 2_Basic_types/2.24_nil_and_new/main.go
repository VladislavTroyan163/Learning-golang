package main

import "fmt"

func main() {
	x := 55

	var p *int = &x
	fmt.Println(p)
	//fmt.Println(*p) - panic

	if p != nil {
		fmt.Println(*p)
	} else {
		fmt.Println("Указатель равен nil")
	}

	pp := new(int)
	fmt.Println(pp)
	fmt.Println(*pp)
	*pp = 50
	fmt.Println(*pp)
}
