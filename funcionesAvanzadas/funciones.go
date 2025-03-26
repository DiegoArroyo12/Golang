package main

import "fmt"

// Función Variádica: Recibe n cantidad de parámetros
func suma(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}

	fmt.Printf("Hola %s, la suma es: %d\n", name, total)

	return total
}

func imprimirDatos(datos ...interface{}){
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

// Función Recursiva: Se ejecuta a sí misma
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}

func main() {
	fmt.Println("Función Variádica")
	fmt.Println(suma("Diego", 12, 45, 78, 56))
	fmt.Println(suma("Alberto", 10, 20, 30, 40, 50))
	imprimirDatos("Hola", 28, true, 3.14)
	fmt.Println("Función Recursiva")
	fmt.Println(factorial(5))
}