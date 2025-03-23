package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello devuelve un saludo para la persona especificada
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("Nombre vacío")
	}

	// Devuelve un saludo que incluye el nombre en un mensaje
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat devuelve uno de varios formatos de mensajes de saludos y se selecciona de forma aleatoria
func randomFormat() string {
	formats := [] string {
		"¡Hola, %v! ¡Bienvenido!",
		"¡Qué bueno verte, %v!",
		"¡Saludos, %v! ¡Encantado de conocerte!",
	}

	return formats[rand.Intn(len(formats))]
}