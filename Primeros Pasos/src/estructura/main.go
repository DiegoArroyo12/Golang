package main

import "fmt"

type Persona struct {
	nombre string
	edad int
	correo string
}

func main() {
	var p Persona
	p.nombre = "Diego"
	p.edad = 21
	p.correo = "dcorreo@gmail.com"

	p2 := Persona {"Juan", 40, "jcorreo@gmail.com"}

	fmt.Println(p)
	fmt.Println(p2.nombre)
}