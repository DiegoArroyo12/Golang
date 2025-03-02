package main

import "fmt"

func main() {
	colors := map[string]string{
		"rojo": "#FF0000",
		"verde": "#00FF00",
		"azul": "#0000FF",
	}

	fmt.Println(colors["rojo"])

	// Agregar un elemento
	colors["negro"] = "000000"
	fmt.Println(colors)

	// Recuperar un valor
	valor, ok := colors["blanco"] // Asigna el valor y en 'ok' marca si existe o no en booleano
	if ok {
		fmt.Println(valor)
	} else {
		fmt.Println("No existe esta clave")
	}

	// Eliminar elemento
	delete(colors, "azul") // De donde y que elemento
	fmt.Println(colors)

	for clave, valor := range colors {
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}
}