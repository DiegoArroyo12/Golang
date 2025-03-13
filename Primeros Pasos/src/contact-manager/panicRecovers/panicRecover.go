package main

import (
	"fmt"
	"log"
	"os"
)

func divideZero(dividendo, divisor int) {
	// Capturamos el pánico para que no se interrumpa la ejecución del código
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No puedes dividir por cero")
	}
}

func main() {
	divideZero(100, 10)
	divideZero(200, 25)
	divideZero(34, 0)
	divideZero(100, 4)

	/*
	log.Print("Este es un mensaje de registro")
	log.SetPrefix("main():") // Establece este prefijo para los errores
	log.Fatal("Oye, soy un registro Fatal y detengo la ejecución")
	log.Panic("Soy un registro de Panico y detengo la ejecución")
	fmt.Println("¿Puedes verme?")
	 */

	// Agregamos un prefijo a los logs
	log.SetPrefix("main(): ")

	file, err := os.OpenFile("info.log", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0644) // Valor octal para permisos de lectura y escritura

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Oye, soy un log")
}