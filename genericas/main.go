package main

import "fmt"

func PrintList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

func main() {
	PrintList("Diego", "Alex", "Emmanuel", "Ricardo")
	PrintList(100, 456, 789, 456, 247)
	PrintList("Hola", 452, 4.25, true)
}