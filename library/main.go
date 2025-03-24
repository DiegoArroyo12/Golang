package main

import "library/book"

func main() {
	var myBook = book.Book{
		Title: "Conde de Montecristo",
		Author: "Alejandro Dumas",
		Pages: 1020,
	}

	secondBook := book.NewBook("La Mano del Muerto", "Alejandro Dumas", 1023)

	myBook.PrintInfo()
	secondBook.PrintInfo()
}