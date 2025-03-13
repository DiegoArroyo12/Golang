package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("no se puede dividir por cero")
		// También se puede hacer con fmt
		// return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return dividendo / divisor, nil
}

func main() {
	// Uso de defer
	file, err := os.Create("hola.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	// Con defer posponemos la ejecución de la instrucción
	// esta se ejecutará al final de esta función 'useDefer'
	defer file.Close()

	_, err = file.Write([]byte("Hola, Diego Arroyo"))

	if err != nil {
		fmt.Println(err)
		return
	}

	// Manejo de Errores
	str := "123"
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Número: ", num)

	result, err := divide(10,0)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Resultado: ", result)
}