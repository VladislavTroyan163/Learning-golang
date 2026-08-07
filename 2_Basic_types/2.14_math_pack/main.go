package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.E)
	fmt.Println(math.Pi)
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sqrt(16))
	fmt.Println(math.Abs(-6))
	//math. Sin, Cos, Tan, Asin, Log, log10, Min, Max и тд

	// Округление
	i := 0.1
	x1 := 0.1 + 0.2
	x2 := i + 0.2
	fmt.Println(x1) // Вывод: 0.3
	fmt.Println(x2) // Вывод: 0.30000000000000004

	fmt.Println(math.Round(-3.4)) //-3 - по правилам математики
	fmt.Println(math.Floor(-3.8)) //-4 - всегда в меньшую сторону
	fmt.Println(math.Floor(-3.2)) //-4 - всегда в меньшую сторону
	fmt.Println(math.Floor(3.8))  // 3 - всегда в меньшую сторону
	fmt.Println(math.Floor(3.2))  // 3 - всегда в меньшую сторону

	fmt.Println(math.Ceil(-3.8)) //-3 - всегда в большую сторону
	fmt.Println(math.Ceil(-3.2)) //-3 - всегда в большую сторону
	fmt.Println(math.Ceil(3.8))  // 4 - всегда в большую сторону
	fmt.Println(math.Ceil(3.2))  // 4 - всегда в большую сторону

	fmt.Println(math.Trunc(-3.8)) //-3 - отбрасывается дробная часть
	fmt.Println(math.Trunc(-3.2)) //-3 - отбрасывается дробная часть
	fmt.Println(math.Trunc(3.8))  // 3 - отбрасывается дробная часть
	fmt.Println(math.Trunc(3.2))  // 3 - отбрасывается дробная часть

	// Округление с точностью

	num := 3.14159254393279
	fmt.Println(num)
	rounded := math.Round(num * 100)
	fmt.Println(rounded / 100)
}

// решение задания:
// fmt.Println("Исходное число:",strconv.FormatFloat(random, 'f', 1, 64))
// fmt.Println("Исходное число, увеличенное на 10%:",strconv.FormatFloat(random * 1.1, 'f', 5, 64))
// fmt.Println("Исходное число является четным:",math.Mod(random, 2) == 0)
// fmt.Println("Предпоследняя цифра целой части исходного числа:",int(random*10)/100)
