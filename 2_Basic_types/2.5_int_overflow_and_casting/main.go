package main

import "fmt"

func main() {
	var num1 int8 = 127
	// var num int8 = 128 - ошибка
	fmt.Println(num1)
	// Но
	num1 = num1 + 1
	fmt.Println(num1) // результат -128 - произошло переполнение (после переполнения счет начинается с начала)

	num2 := 200
	fmt.Println(num2)
	var num3 int64
	//num3 = num2 - ошибка, так как int и int64 считаются разными типами данных
	num3 = int64(num2) // int представлен как int64 - ошибки нет
	fmt.Println(num3)
	//num1 = num3 - ошибка, так как int8 и int64 считаются разными типами данных
	num1 = int8(num3)
	fmt.Println(num1) // Из за переполнения результат -56
}

// Тестирование на stepik:
// [+] Test #1. OK
// [+] Test #2. OK

// 2 of 2 test(s) passed.

// // переменная с названием i8 уже создана, в ней лежит значение с типом данных int8
// var i int
// i = int(i8)
