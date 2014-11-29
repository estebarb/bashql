package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"os"
	"unicode/utf8"
)

var separator = flag.String("d", ",", "field delimiter (',' by default)")
var comment = flag.String("c", "", "help message for flagname")

func main() {
	flag.Parse()

	// Abre el archivo para ser leído
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	// Configura la lectura del CSV
	r := csv.NewReader(reader)
	r.Comma, _ = utf8.DecodeRuneInString(*separator)
	if *comment != "" {
		r.Comment, _ = utf8.DecodeRuneInString(*comment)
	}

	// Abrimos el nuevo CSV en la salida estandar
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	w := csv.NewWriter(writer)
	defer w.Flush()

	// Imprime el contenido del archivo
	col, er := r.Read()
	for er == nil {
		w.Write(col)
		col, er = r.Read()
	}
}
