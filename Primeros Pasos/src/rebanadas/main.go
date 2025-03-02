package main

import "fmt"

func main()  {
	diasSemana := [] string{"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}
	diaRebanada := diasSemana[0:5]
	fmt.Println(diasSemana)

	// Agregar elementos al slice
	diaRebanada = append(diaRebanada, "Viernes", "Sábado", "Otro día") // Se aumenta la capacidad del slice

	fmt.Println(len(diaRebanada))
	fmt.Println(cap(diaRebanada))

	// Crear un slice a partir de otro
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)
	fmt.Println(diaRebanada)

	nombres := make([]string, 5, 10) // Que datos se agregaran, longitud y la capacidad
	nombres[0] = "Diego" // Después de agregar los 5 elementos debemos ocupar append
	fmt.Println(nombres)

	rebanada := []int{1,2,3,4,5}
	rebanadaDos := make([] int, 5)

	copy(rebanadaDos, rebanada) // Los elementos de rebanada se copian a rebanadaDos

	fmt.Println(copy(rebanadaDos, rebanada)) // La cantidad de los elementos copiados en este caso 5

	fmt.Println(rebanada)
	fmt.Println(rebanadaDos)
}