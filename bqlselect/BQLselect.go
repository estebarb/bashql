package main

import (
	"bufio"
	"encoding/csv"
	"github.com/estebarb/bashql/common"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	r := csv.NewReader(reader)
	// La primera fila indica el nombre de las columnas,
	columnas, err := r.Read()

	if err != nil {
		panic(err)
	}

	// Los argumentos del programa corresponden con
	// las columnas seleccionadas:
	argumentos := os.Args[1:]
	indices, err := common.Filtrados(argumentos, columnas)
	if err != nil {
		panic(err)
	}

	// Abrimos el nuevo CSV en la salida estandar
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	w := csv.NewWriter(writer)
	defer w.Flush()

	// Imprimimos las cabeceras:
	w.Write(common.Seleccionar(indices, columnas))

	col, er := r.Read()
	for er == nil {
		w.Write(common.Seleccionar(indices, col))
		col, er = r.Read()
	}
}
