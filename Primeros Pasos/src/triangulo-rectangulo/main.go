package main

import (
	"fmt"
	"math"
)

const precision = 2

func main(){
	var ladoUno, ladoDos float64

	// Entrada de Datos
	fmt.Print("Ingrese el lado 1: ")
	fmt.Scanln(&ladoUno)
	fmt.Print("Ingrese el lado 2: ")
	fmt.Scanln(&ladoDos)

	// Procesos
	hipotenusa := math.Sqrt(math.Pow(ladoUno, 2) + math.Pow(ladoDos, 2))
	area := (ladoUno * ladoDos) / 2
	perimetro := ladoUno + ladoDos + hipotenusa

	fmt.Printf("\nÁrea: %.*f", precision, area) // Limitamos a 2 decimales
	fmt.Printf("\nPerímetro: %.*f", precision, perimetro)
}