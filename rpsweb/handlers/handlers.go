package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error al analizar plantillas", http.StatusInternalServerError)
		return
	}

	data := struct {
		Title	string
		Message	string
	} {
		Title: "Página de Inicio",
		Message: "¡Bienvenido a Piedra, Papel o Tijera!",
	}

	err = tpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		return
	}
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear Nuevo Juego")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Juego")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Acerca de")
}