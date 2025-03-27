package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

type integer int 

type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

// Restricción arbitraria
func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list[]T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}

type Product[T uint | string] struct {
	Id T
	Desc string
	Price float32
}

func main() {
	PrintList("Diego", "Alex", "Emmanuel", "Ricardo")
	PrintList(100, 456, 789, 456, 247)
	PrintList("Hola", 452, 4.25, true)

	fmt.Println(Sum[uint](4, 5, 6, 4))
	fmt.Println(Sum[float32](4.5, 5.5, 6.5, 4.5))

	var num1 integer = 100
	var num2 integer = 300

	fmt.Println(Sum(num1, num2))

	strings := []string{"a", "b", "c", "d"}
	numbers := []int{1,2,3,4,5,6}

	fmt.Println(Includes(strings, "a"))
	fmt.Println(Includes(strings, "f"))
	fmt.Println(Includes(numbers, 4))
	fmt.Println(Includes(numbers, 5))

	fmt.Println(Filter(numbers, func(value int) bool {return value > 3}))
	fmt.Println(Filter(strings, func(value string) bool {return value > "b"}))

	product1 := Product[uint]{1, "Zapatos", 50}
	product2 := Product[string]{"SDA-ASDG8-ASSASSA", "Zapatos", 50}
	fmt.Println(product1, product2)
}