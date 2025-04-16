package testunitario

import "testing"

/* func TestSuma(t *testing.T) {
	total := Suma(5, 5)

	if total != 11 {
		t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, 11)
	}
} */
/* 
go test

--- FAIL: TestSuma (0.00s)
    mate_test.go:9: Suma incorrecta, tiene 10 se esperaba 11
FAIL
exit status 1
FAIL    github.com/DiegoArroyo12/testing/testunitario   0.229s
*/

func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	} {
		{1,2,3},
		{2,2,4},
		{25,25,50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("Suma incorrecta, tiene %d se esperaba %d", total, item.c)
		}
	}
}