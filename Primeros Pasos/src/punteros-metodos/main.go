package main

import "fmt"

type Persona struct {
	nombre string
	edad int
	correo string
}

func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es", p.nombre)
}

func main() {
	var x int = 10
	fmt.Println(x)
	//var p *int = &x // Asignamos un puntero a x

	editar(&x)
	fmt.Println(x)

	p := Persona{"Diego", 21, "dcorreo@gmail.com"}
	p.diHola()
}

func editar(x *int) {
	*x = 20
}