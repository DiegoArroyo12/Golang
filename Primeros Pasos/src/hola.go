package main

import (
	"fmt"
	"rsc.io/quote"
)

// Una forma más simplificada de declarar variables
var nombre, apellido, edad = "Diego", "Arroyo", 21

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
}