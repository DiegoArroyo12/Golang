package greetings

import (
	"errors"
	"fmt"
)

// Hello devuelve un saludo para la persona especificada
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Nombre vacío")
	}

	// Devuelve un saludo que incluye el nombre en un mensaje
	message := fmt.Sprintf("¡Hola, %v! ¡Bienvenido!", name)
	return message, nil
}