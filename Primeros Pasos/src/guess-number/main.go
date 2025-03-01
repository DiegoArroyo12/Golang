package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	t := time.Now()
	hora := t.Hour()

	if hora < 12 {
		fmt.Println("¡Mañana!")
	} else if hora < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}

	// Otra forma de hacer lo mismo (t solo se puede usar dentro del if)
	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("¡Mañana!")
	} else if t.Hour() < 17 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}

	// Switch
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Se esta ejecutando en Windows")
	case "linux":
		fmt.Println("Se esta ejecutando en Linux")
	case "darwin":
		fmt.Println("Se esta ejecutando en MacOS")
	default:
		fmt.Println("Se ejecuta en otro Sistema Operativo")
	}

	// Ejercicio del tiempo con switch
	switch t := time.Now(); {
		case t.Hour() < 12:
			fmt.Println("¡Mañana!")
		case t.Hour() < 17:
			fmt.Println("¡Tarde!")
		default:
			fmt.Println("¡Noche!")
	}

	// For
	var i int

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}
}