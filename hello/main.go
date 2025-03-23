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

	names := []string{"Diego", "David", "Denisse"}
	messages, err := greetings.Hellos((names))

	/* message, err := greetings.Hello("Diego")

	if err != nil {
		log.Fatal(err)
	} */

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}