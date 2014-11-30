package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"os"
)

var insert bool

func init() {
	flag.BoolVar(&insert, "i", false, "Inserts new header")
	flag.BoolVar(&insert, "insert", false, "Inserts new header")
}

func main() {
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	r := csv.NewReader(reader)

	// Abrimos el nuevo CSV en la salida estandar
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	w := csv.NewWriter(writer)
	defer w.Flush()

	// Imprimimos las cabeceras:
	if !insert {
		_, err := r.Read()
		if err != nil {
			panic(err)
		}
	}
	w.Write(flag.Args())

	col, er := r.Read()
	for er == nil {
		w.Write(col)
		col, er = r.Read()
	}
}
