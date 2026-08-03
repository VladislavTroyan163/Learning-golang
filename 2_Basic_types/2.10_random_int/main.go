package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func main() {
	// randimNum := rand.IntN(100) // [0, 100)
	// fmt.Println(randimNum)

	// Псевдо-случайные числа
	// min := 10
	// max := 50
	// randimNum := rand.IntN(max - min) // [0, 40)
	// fmt.Println(randimNum + min)      // [10, 50)

	// Супер пупер случайные числа
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		log.Fatalf("Ошибка генерации случайного числа %v", n.Int64())
	}
	fmt.Println("Случайное число:", n.Int64())
}
