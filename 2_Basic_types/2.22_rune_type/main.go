package main

import "fmt"

func main() {
	// руна - это символ
	var a1 rune = 'A' // = int32 //Присвоение одиночными кавычками
	fmt.Printf("Индекс: %d\nСимвол: %c\n", a1, a1)

	var a2 rune = '\u1234'
	fmt.Printf("Индекс: %d\nСимвол: %c\n", a2, a2)

	a3 := rune(128512)
	fmt.Printf("Индекс: %d\nСимвол: %c\n", a3, a3)

	str1 := "Hello"
	fmt.Printf("Индекс: %d\nСимвол: %c\n", str1[0], str1[0])
	fmt.Printf("Индекс: %d\nСимвол: %c\n", str1[1], str1[1])

	str2 := "Привет"
	fmt.Printf("Индекс: %d\nСимвол: %c\n", str2[0], str2[0])
	fmt.Printf("Индекс: %d\nСимвол: %c\n", str2[1], str2[1])
	// Так как русские буквы в алфавите занимают 2 байта, выводятся другое символы

	str3 := "Привет"
	runes := []rune(str3)
	fmt.Println(runes)
	fmt.Printf("Индекс: %d\nСимвол: %c\n", runes[0], runes[0])
}
