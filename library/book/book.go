package book

import "fmt"

type Book struct {
	Title string
	Author string
	Pages int
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title: title,
		Author: author,
		Pages: pages,
	}
}

func (b *Book) PrintInfo(){
	fmt.Printf("Title: %s\nAuthor: %s\nPages: %d\n", b.Title, b.Author, b.Pages)
}