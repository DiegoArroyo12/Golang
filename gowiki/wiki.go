package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600) // Persmisos de lectura y escritura
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("Esta es una página de muestra")}
	p1.save()

	p2, _ := loadPage("TestPage")
	fmt.Println(p2.Body) // En bytes
	fmt.Println(string(p2.Body)) // En string
}