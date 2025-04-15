package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK		= 0 // Piedra. Vence a tijeras. (tijeras + 1) % 3 = 0
	PAPER		= 1 // Papel. Vence a la piedra. (piedra + 1) % 3 = 1
	SCISSORS	= 2 // Tijeras. Vence al papel. (papel + 1) % 3 = 2
)

// Estructura para dar resultado de cada ronda
type Round struct {
	Message				string 	`json:"message"`
	ComputerChoice		string 	`json:"computer_choice"`
	RoundResult			string 	`json:"round_result"`
	ComputerChoiceInt	int 	`json:"computer_choice_int"`
	ComputerScore		string	`json:"computer_score"`
	PlayerScore			string	`json:"player_score"`
}

// Mensajes para cuando gana
var winMessages = []string {
	"¡Bien Hecho!",
	"¡Buen Trabajo!",
	"Deberías comprar un boleto de lotería",
}

// Mensajes para cuando pierde
var loseMessages = []string {
	"¡Qué Lástima!",
	"¡Inténtalo de Nuevo!",
	"Hoy simplemente no es tu día",
}

// Mensaje de empate
var drawMessages = []string {
	"Las grandes mentes piensan igual.",
	"Oh oh. Inténtalo de nuevo.",
	"Nadie gana, pero puedes intentarlo de nuevo.",
}

// Variables para el puntaje
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	// Mensaje dependiendo de lo que eligió la computadora
	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligió PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligió PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligió TIJERAS"
	}

	// Generar un número aleatorio de 0-2, que usamos para elegir un mensaje aleatorio
	messageInt := rand.Intn(3)

	// Declarar una var para contener el mensaje
	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		// Seleccionar mensaje de drawMessages
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue + 1) {
		PlayerScore++
		roundResult = "¡El jugador gana!"
		// Seleccionar mensaje de winMessages
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "¡La computadora gana!"
		// Seleccionar mensaje de loseMessages
		message = loseMessages[messageInt]
	}

	// Retornar resultado
	return Round{
		Message: message,
		ComputerChoice: computerChoice,
		RoundResult: roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore: strconv.Itoa(ComputerScore),
		PlayerScore: strconv.Itoa(PlayerScore),
	}
}