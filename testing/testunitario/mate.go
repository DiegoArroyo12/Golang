package testunitario

func Suma(a, b int) int {
	return a + b
}

func GetMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}

/* 
go test -cover (Porcentaje testeado)
PASS
coverage: 25.0% of statements
ok      github.com/DiegoArroyo12/testing/testunitario   0.208s


go test -coverprofile=coverage.out (Creación de archivo .out)
PASS
coverage: 25.0% of statements
ok      github.com/DiegoArroyo12/testing/testunitario   0.204s


go tool cover -func=coverage.out (Muestra más legible del archivo)
github.com/DiegoArroyo12/testing/testunitario/mate.go:3:        Suma            100.0%
github.com/DiegoArroyo12/testing/testunitario/mate.go:7:        GetMax          0.0%
total:                                                          (statements)    25.0%


go tool cover -html=coverage.out (Abre en el navegador una forma más visual de ver lo testeado)
*/

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n - 1) + Fibonacci(n - 2)
}