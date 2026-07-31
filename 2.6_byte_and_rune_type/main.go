package main

import "fmt"

func main() {
	//byte - тип, часто используется для ASCII, аналог uint8
	//uint8 - это число, а когда пишем byte - подразумеваем, что это определенный байт от0 до 255
	var b byte = 200
	fmt.Println(b, string(b))

	s := "Hi"
	fmt.Println(s[0], string(s[0]))
	fmt.Println(s[1], string(s[1]))

	//rune - аналог int32
	//int32 - это число, а когда пишем rune - подразумеваем, что это символ
	var r rune
	r = '💖'
	fmt.Println(r, string(r))
}

//Это всё сделано для уобства работы с unicode
