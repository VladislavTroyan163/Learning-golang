package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//float32 and float64
	fmt.Printf("max float32: %v\n", math.MaxFloat32)
	fmt.Printf("max float64: %v\n", math.MaxFloat64)

	var a float32
	fmt.Println(a, reflect.TypeOf(a))

	var b float32 = 34_553.1_415_756_537_875
	fmt.Println(b, reflect.TypeOf(b))

	c := 3.14154
	fmt.Println(c, reflect.TypeOf(c))

	// NaN - не число
	// Inf - бесконечность

	// res1 := 1/0 // ошибка
	// zero1 := 0
	// res1 := 1 / zero1
	// fmt.Println(res1)

	zero2 := 0.0
	res2 := 1.0 / zero2
	fmt.Println(res2)

	res3 := zero2 / zero2
	fmt.Println(res3)

	nan := math.NaN()
	posInf := math.Inf(1)
	negInf := math.Inf(-1)
	fmt.Println("NaN?", math.IsNaN(nan))
	fmt.Println("NaN?", math.IsInf(posInf, 1))
	fmt.Println("NaN?", math.IsInf(negInf, -1))
	fmt.Println("NaN?", math.IsInf(posInf, 0))
	fmt.Println("NaN?", math.IsInf(negInf, -0))

	i := 42
	var f float64 = float64(i)
	fmt.Println(f)

	f1 := 55.5
	var i1 int = int(f1)
	var i2 int = int(math.Round(f1))
	fmt.Println(i1)
	fmt.Println(i2)

}
