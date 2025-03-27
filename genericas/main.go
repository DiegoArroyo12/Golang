package main

import "fmt"

func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

type integer int 

// Restricción arbitraria
func Sum[T ~int | ~float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	PrintList("Diego", "Alex", "Emmanuel", "Ricardo")
	PrintList(100, 456, 789, 456, 247)
	PrintList("Hola", 452, 4.25, true)

	fmt.Println(Sum[int](4, 5, 6, 4))
	fmt.Println(Sum(4.5, 5.5, 6.5, 4.5))

	var num1 integer = 100
	var num2 integer = 300

	fmt.Println(Sum(num1, num2))
}