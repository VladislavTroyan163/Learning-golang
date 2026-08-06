package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "12345"
	// num, err := strconv.Atoi(str)
	// if err != nil {
	// 	fmt.Println("ОШИБКА!!!")
	// } else {
	// 	fmt.Println(num)
	// }

	num, err := strconv.ParseInt(str, 10, 64)
	//(str, 0, 64) - автоматическое определение СС входной строки
	//(str, 10, 0) - стандартный размер типа int
	if err != nil {
		fmt.Println("ОШИБКА!!!")
	} else {
		fmt.Println(num)
	}

	str2 := "3.25"
	num2, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		fmt.Println("ОШИБКА!!!")
	} else {
		fmt.Println(num2)
	}

}
