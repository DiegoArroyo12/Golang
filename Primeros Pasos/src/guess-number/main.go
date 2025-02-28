package main

import (
	"fmt"
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
}