package main

import (
	"fmt"
)

func main() {
	// var firstname string
	// // var lastname string
	// fmt.Println("Введите имя:")
	// fmt.Scanln(&firstname)
	// // fmt.Println("Введите фамилию:")
	// // fmt.Scanln(&lastname)
	// // fmt.Printf("Привет, %s %s, ля ля ля\n", firstname, lastname)

	// // fmt.Println("Введите имя и фамилию через пробел:")
	// // fmt.Scanln(&firstname, lastname)
	// // fmt.Printf("Привет, %s %s, ля ля ля\n", firstname, lastname)
	// fmt.Printf("Привет, %s\n", firstname)

	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("Введите строку: ")
	// scanner.Scan()
	// fmt.Println("Вы ввели:", scanner.Text())
	inputPeopleInfo()
}

// Задание для самостоятельной работы:
func inputPeopleInfo() {
	var firstname string
	var lastname string
	var age int
	fmt.Print("Введите имя и фамилию через пробел:")
	fmt.Scanln(&firstname, &lastname)
	fmt.Print("Введите возраст:")
	fmt.Scanln(&age)
	fmt.Printf("Приятно познакомиться, %s. Я 5 лет назад познакомился с человеком, у которого тоже фамилия %s, вам тогда было %d. Как молоды мы были!", firstname, lastname, age-5)
}
