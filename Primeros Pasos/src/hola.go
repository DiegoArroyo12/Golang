package main

import (
	"fmt"
	"rsc.io/quote"
)

// Una forma más simplificada de declarar variables
var nombre, apellido, edad = "Diego", "Arroyo", 21

// Declaración de Constantes
const Pi = 3.14

const (
	x = 100
	y = 0b1010	// Binario
	z = 0o12	// Octal
	w = 0xFF	// Hexadecimal
)

const (
	Domingo = iota + 1	// Va aumentando el valor para las siguientes constantes
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado 
)

func main(){
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Hello())

	// Declaración de variables
	/* Forma básica de hacerlo
	var firstName, lastName string
	var age int 
	*/

	var (
		// Si no inicializamos el valor de la variable, es necesario poner el tipo
		//firstName, lastName string
		//age int = 27
		firstName = "Diego"
		lastName = "Arroyo"
		age = 21
	)	

	// Declaración sin la palabra reservada 'var' que solo funciona dentro de una función
	lenguaje, tiempo, computadora := "Golang", 18, "Mac"

	/* 
	firstName = "Diego"
	lastName = "Arroyo"
	age = 21
	*/

	fmt.Println(firstName, lastName, age)
	fmt.Println(nombre, apellido, edad)
	fmt.Println(lenguaje, tiempo, computadora)
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)
}