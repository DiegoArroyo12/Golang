package main

import (
	"html/template"
	"log"
	"net/http"
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

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)

	if err != nil {
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	//fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body) // Capturar después de la '/'
	renderTemplates(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)

	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplates(w, "edit", p)
}

func renderTemplates(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func main() {
	//p1 := &Page{Title: "TestPage", Body: []byte("Esta es una página de muestra")}
	//p1.save()

	//p2, _ := loadPage("TestPage")
	//fmt.Println(p2.Body) // En bytes
	//fmt.Println(string(p2.Body)) // En string

	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}