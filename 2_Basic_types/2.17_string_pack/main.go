package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "  Привет Алексей!!!"
	fmt.Println(str)

	fmt.Println(strings.TrimSpace(str)) // убрать пробелы до и после

	fmt.Println(strings.ToUpper(str)) // сделать капс
	fmt.Println(strings.ToLower(str)) // сделать !капс

	words := strings.Split(str, " ") // Разбить на слова по признаку пробела (получаем слайс)
	fmt.Println(words)

	fmt.Println(strings.Contains(str, "При"))
	fmt.Println(strings.Contains(str, "при")) // регистрозависимость

	fmt.Println(strings.HasPrefix(str, "  П"))
	fmt.Println(strings.HasPrefix(str, "П"))
	fmt.Println(strings.HasSuffix(str, "!"))
	fmt.Println(strings.HasSuffix(str, "!?"))

	fmt.Println(strings.EqualFold("ПУПА и ЛуПа", "Пупа И ЛУПА")) // Сравнение строк без учера регистра
}
