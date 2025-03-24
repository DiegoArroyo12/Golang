package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {
	/* var myBook = book.Book{
		Title: "Conde de Montecristo",
		Author: "Alejandro Dumas",
		Pages: 1020,
	} */

	myBook := book.NewBook("La Mano del Muerto", "Alejandro Dumas", 1023)
	myBook.PrintInfo()

	myBook.SetTitle("Conde de Montecristo")
	fmt.Println(myBook.GetTitle())

	myTextBook := book.NewTexBook("Programación", "Jaime Gamarra", 261, "Santillana SAC", "Secundaria")
	myTextBook.PrintInfo()

	book.Print(myBook)
	book.Print(myTextBook)

	// Objetos de Animal
	miPerro := animal.Perro{Nombre: "Rocko"}
	miGato := animal.Gato{Nombre: "Chichara"}

	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

	animales := []animal.Animal {
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Tom"},
		&animal.Perro{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Luna"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}