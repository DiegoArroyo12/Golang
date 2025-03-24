package book

import "fmt"

type Book struct {
	// En minúsculas se hacen atributos privados
	title string
	author string
	pages int
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title: title,
		author: author,
		pages: pages,
	}
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo(){
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}