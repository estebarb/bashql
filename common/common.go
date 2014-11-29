package common

import (
	"fmt"
	"strconv"
)

// Devuelve los índices de las columnas seleccionadas en
// los argumentos, según el índice en las columnas.
func Filtrados(argumentos []string, columnas []string) ([]int, error) {
	salida := make([]int, len(argumentos))

	for k, v := range argumentos {
		encontrado := false
		for k2, v2 := range columnas {
			if v == v2 {
				encontrado = true
				salida[k] = k2
				break
			}
		}

		// Como no se encontró puede ser que el argumento
		// fuera el número de columna
		if !encontrado {
			num, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return nil, err
			}
			if num < 0 || int(num) > len(columnas) {
				return nil, fmt.Errorf("Column number %d is out of range",
					num)
			}
			salida[k] = int(num)
		}
	}

	return salida, nil
}

// Devuelve las columnas seleccionadas
func Seleccionar(seleccionados []int, entradas []string) []string {
	salida := make([]string, len(seleccionados))
	for k, v := range seleccionados {
		salida[k] = entradas[v]
	}
	return salida
}

func GetValue(col string, datos, columns []string) string{
	for k, v := range(columns){
		if v == col {
			return datos[k]
		}
	}
	panic(fmt.Errorf("%v column not found", col))
}
