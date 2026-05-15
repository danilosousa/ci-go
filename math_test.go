package main

import "testing"

func TestSoma(t *testing.T) {
	result := soma(20, 2)
	if result != 22 {
		t.Errorf("o resultado da soma é invalido = %d; o esperado %d", result, 22)
	}
}
