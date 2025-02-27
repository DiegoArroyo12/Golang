package main

import (
	"fmt"
	"strconv"

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

	// Valores Predeterminados
	var (
		defaultInt int
		defautlUint uint
		defaultFloat float32
		defaultBool bool
		defaultString string
	)

	fmt.Println(defaultInt, defautlUint, defaultFloat, defaultBool, defaultString) // 0 0 0 false '' (sin comillas, cadena vacía)

	fmt.Println(firstName, lastName, age)
	fmt.Println(nombre, apellido, edad)
	fmt.Println(lenguaje, tiempo, computadora)
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Viernes)

	// Conversiones de tipos
	var integer16 int16 = 50
	var integer32 int32 = 100

	// De integer32 a integer16
	fmt.Println(integer16 + int16(integer32))

	s := "100"
	i, _ := strconv.Atoi(s) // El guión bajo es para controlar el error

	// De string a int
	fmt.Println(i + i)

	n := 42
	s = strconv.Itoa(n)

	// De int a string
	fmt.Println(s + s)

	// Printf
	name := "Diego"
	ages := 21

	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, ages)

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.\n", name, ages)
	fmt.Println(greeting)

	// Ingresar usuarios en terminal
	var nombreUsuario string
	var apellidoUsuario string
	var edadUsuario int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombreUsuario, &apellidoUsuario)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&edadUsuario)

	 // %v cuando no sabemos que tipo de dato se va a mostrar
	 // %T es para saber el tipo de dato con printf
	fmt.Printf("Hola, me llamo %s %s y tengo %d años.\n", nombreUsuario, apellidoUsuario, edadUsuario)
}