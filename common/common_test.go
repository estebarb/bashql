package common

import (
	"testing"
)

func TestFiltrados(t *testing.T) {
	deseados := []string{"tercera", "segunda"}
	columnas := []string{"primera", "segunda", "tercera"}
	esperado := []int{2, 1}
	out, err := Filtrados(deseados, columnas)
	if err != nil {
		t.Error(err)
	}
	if !compararInt(out, esperado) {
		t.Errorf("Filtrados(%v,%v) = %v, want %v", deseados, columnas, out, esperado)
	}
}

func TestSeleccionar(t *testing.T) {
	seleccionar := []int{2, 1}
	columnas := []string{"primera", "segunda", "tercera"}
	esperados := []string{"tercera", "segunda"}
	salida := Seleccionar(seleccionar, columnas)
	if !compararString(salida, esperados) {
		t.Errorf("Filtrados(%v,%v) = %v, want %v", seleccionar, columnas, salida, esperados)
	}
}

func compararString(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}

func compararInt(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}
