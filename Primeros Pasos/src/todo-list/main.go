package main

import "fmt"

func main() {
	var a [5] int

	a[0] = 10
	a[1] = 20

	var b = [...] int{10, 20, 30, 40, 50} // Los tres puntos es cuando no sabemos cuantos elementos tendrá

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b[2])

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	// Manejo de rango
	for index, value := range b {
		fmt.Printf("Índice:  %d, Valor: %d\n", index, value)
	}

	// Matriz [filas][columnas]
	var matriz = [3][3] int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz) // Matriz bidimensional 
}