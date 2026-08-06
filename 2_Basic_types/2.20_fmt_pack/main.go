package main

import "fmt"

func main() {
	fmt.Print("Привет, ")
	fmt.Print("golang\n")

	name := "Владислав"
	fmt.Println("Привет, ", name)
	fmt.Print("Привет, ", name)

	fmt.Printf("\nЛЯ ля ля %s %s", name, "Троян\n")

	age := 24
	str := fmt.Sprint("Меня зовут ", name, ", мне ", age, " года")
	fmt.Println(str)

	str = fmt.Sprintf("Меня зовут %s, мне %d года", name, age)
	fmt.Println(str)
}
