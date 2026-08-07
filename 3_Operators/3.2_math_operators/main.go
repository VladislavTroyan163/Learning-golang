package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

func main() {
	a := 10
	b := 10.5
	fmt.Println(a + 3)
	fmt.Println(a - 3)
	fmt.Println(a * 3)
	fmt.Println(a / 3)
	fmt.Println(a % 3) // Остаток от деления

	fmt.Println(a % 2)  // Проверка на четность
	fmt.Println(a % 10) // Получить последнюю цифру

	fmt.Println(b + 3)
	fmt.Println(b - 3)
	fmt.Println(b * 3)
	fmt.Println(b / 3)

	fmt.Println(b - 3*math.Trunc(b/3)) // Остаток от деления
	fmt.Println(math.Mod(b, 3))        // Остаток от деления

	num := 5
	fmt.Println(num)
	num++
	fmt.Println(num)
	num--
	fmt.Println(num)

	random := math.Floor((rand.Float64()*100)*10) / 10
	fmt.Println(random)

}
