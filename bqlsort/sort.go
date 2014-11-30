package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"github.com/estebarb/bashql/common"
	"os"
	"sort"
)

var reverse bool

func init() {
	flag.BoolVar(&reverse, "r", false, "Use reverse order")
	flag.BoolVar(&reverse, "reverse", false, "Use reverse order")
}

func main() {
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	r := csv.NewReader(reader)
	// La primera fila indica el nombre de las columnas,
	columnas, err := r.Read()
	if err != nil {
		panic(err)
	}

	// Abrimos el nuevo CSV en la salida estandar
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	w := csv.NewWriter(writer)
	defer w.Flush()

	// Imprimimos las cabeceras:
	w.Write(columnas)

	// cargamos todos los datos
	rawData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	data := common.GenMultiValueTable(columnas, flag.Args(), rawData)

	if !reverse {
		sort.Sort(common.ByMultiValue(data))
	} else {
		sort.Sort(sort.Reverse(common.ByMultiValue(data)))
	}
	for _, row := range data {
		w.Write(row.Data)
	}
}
