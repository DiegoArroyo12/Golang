package main

import (
	"fmt"

	"github.com/DiegoArroyo12/greetings"
)

func main() {
	message := greetings.Hello("Diego")
	fmt.Println(message)
}