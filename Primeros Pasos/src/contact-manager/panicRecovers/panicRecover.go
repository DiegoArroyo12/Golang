package main

import "fmt"

func divideZero(dividendo, divisor int) {
	// Capturamos el pánico para que no se interrumpa la ejecución del código
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No puedes dividir por cero")
	}
}

func main() {
	divideZero(100, 10)
	divideZero(200, 25)
	divideZero(34, 0)
	divideZero(100, 4)
}