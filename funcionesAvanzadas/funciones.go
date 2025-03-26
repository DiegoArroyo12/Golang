package main

import "fmt"

// Función Variádica: Recibe n cantidad de parámetros
func suma(name string, nums ...int) int {
	var total int
	for _, num := range nums {
		total += num
	}

	fmt.Printf("Hola %s, la suma es: %d\n", name, total)

	return total
}

func imprimirDatos(datos ...interface{}){
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

// Función Recursiva: Se ejecuta a sí misma
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}

// Función Anónima
func saludar(name string, f func(string))  {
	f(name)
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

// Funciones de Orden Superior
func doble(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}

// Clousures: Esta hace que la función anónima recuerde el valor de i en este ejemplo
func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	fmt.Println("Función Variádica")
	fmt.Println(suma("Diego", 12, 45, 78, 56))
	fmt.Println(suma("Alberto", 10, 20, 30, 40, 50))
	imprimirDatos("Hola", 28, true, 3.14)

	fmt.Println("Función Recursiva")
	fmt.Println(factorial(5))
	// Función Anónima: Función sin nombre
	fmt.Println("Función Anónima")
	func (){
		fmt.Println("Hola soy una función anónima")
	}()

	saludo := func (name string) {
		fmt.Printf("Hola %s, soy una función anónima\n", name)
	}

	saludo("Diego")
	saludar("David", saludo)

	r1 := aplicar(duplicar, 5)
	r2 := aplicar(triplicar, 5)
	fmt.Println(r1, r2)

	fmt.Println("Función de Orden Superior")
	r := doble(addOne, 3)
	fmt.Println("Resultado: ", r)

	fmt.Println("Closures")
	nextInt := incrementer()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}