package main

import (
	"fmt"
	"log"

	"github.com/DiegoArroyo12/greetings"
)

func main() {
	// Agregamos el prefijo a los errores y ocultamos fecha y hora del error
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}