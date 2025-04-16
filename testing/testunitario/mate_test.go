package testunitario

import "testing"

func TestSuma(t *testing.T) {
	total := Suma(5, 5)

	if total != 11 {
		t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, 11)
	}
}
/* 
go test

--- FAIL: TestSuma (0.00s)
    mate_test.go:9: Suma incorrecta, tiene 10 se esperaba 11
FAIL
exit status 1
FAIL    github.com/DiegoArroyo12/testing/testunitario   0.229s
*/