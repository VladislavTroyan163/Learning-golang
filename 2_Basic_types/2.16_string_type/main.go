package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// строка иммутабельна
	var str1 string
	fmt.Println(str1)
	var str2 = "Привет!"
	fmt.Println(str2)
	str3 := "Ещё \nпривет"
	fmt.Println(str3)
	str4 := `Ещё \n при
		вет"`
	fmt.Println(str4)

	str5 := "Hello"
	fmt.Println(str5, len(str5)) // len берёт количество байт
	str6 := "Привет"
	fmt.Println(str6, len(str6))                                            // len берёт количество байт
	fmt.Println(utf8.RuneCountInString(str5), utf8.RuneCountInString(str6)) //берёт количество символов

	var str7 string = "Привет, \n\"друг\", \tкак \tтвои \tдела😎? 5\\5 = 1"
	fmt.Println(str7)

}
