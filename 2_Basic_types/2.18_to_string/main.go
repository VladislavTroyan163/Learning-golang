package main

import (
	"fmt"
	"strconv"
)

func main() {

	i := 555
	var str string
	str = string(i) // Вывод числа в виде строки (приведение к символу ȫ)
	fmt.Println(str)

	str1 := strconv.Itoa(i)
	fmt.Println(str1)

	str2 := strconv.FormatInt(int64(i), 2) // указание системы счисления
	fmt.Println(str2)

	str3 := strconv.FormatInt(int64(i), 8) // указание системы счисления
	fmt.Println(str3)

	num := 10.3443
	fixed := strconv.FormatFloat(num, 'f', 5, 64)
	fmt.Println("Фиксирорванное представление (f):", fixed)

	num2 := 10.3443
	fixed2 := strconv.FormatFloat(num2, 'e', 2, 64)
	fmt.Println("Фиксирорванное представление (e):", fixed2)

	str4 := fmt.Sprintf("Привет, вот число: %f", num)
	fmt.Println(str4)
}
