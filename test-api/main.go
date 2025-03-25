package main

import (
	"fmt"
	"net/http"
	"time"
)

func main()  {
	start := time.Now()

	apis := []string {
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	for _, api := range apis {
		// La palabra go, inicia la concurrencia
		go checkAPI(api)
	}

	time.Sleep(5 * time.Second)

	elapsed := time.Since(start)
	fmt.Printf("¡Listo! ¡Tomó %v segundos!\n", elapsed.Seconds())
}

func checkAPI(api string) {
	if _, err := http.Get(api); err != nil {
		fmt.Printf("Error: ¡%s está caído!\n", api)
		return
	}

	fmt.Printf("Success: ¡%s está en funcionamiento!\n", api)
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