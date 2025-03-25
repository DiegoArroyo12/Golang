package main

import (
	"fmt"
	"net/http"
	"time"
)

func main()  {
	/* // Declaración de canal
	canal := make(chan int)
	// Enviar datos
	canal <- 15
	// Recibir datos del canal
	valor := <- canal */

	start := time.Now()

	apis := []string {
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	ch := make(chan string)
	for _, api := range apis {
		// La palabra go, inicia la concurrencia
		go checkAPI(api, ch)
	}

	for i := 0; i < len(apis); i++ {
		fmt.Print(<- ch)
	}

	time.Sleep(5 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomó %v segundos!\n", elapsed.Seconds())
}

func checkAPI(api string, ch chan string) {
	if _, err := http.Get(api); err != nil {
		// Enviamos al canal
		ch <- fmt.Sprintf("Error: ¡%s está caído!\n", api)
		return
	}

	ch <- fmt.Sprintf("Success: ¡%s está en funcionamiento!\n", api)
}

/* 
Sin Concurrencia

Success: ¡https://management.azure.com está en funcionamiento!
Success: ¡https://dev.azure.com está en funcionamiento!
Success: ¡https://api.github.com está en funcionamiento!
Success: ¡https://outlook.office.com/ está en funcionamiento!
Error: ¡https://api.somewhereintheinternet.com/ está caído!
Success: ¡https://graph.microsoft.com está en funcionamiento!
¡Listo! ¡Tomó 3.962020375 segundos!

Con Concurrencia

¡Listo! ¡Tomó 3.5708e-05 segundos!
*/