package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/estebarb/bashql/common"
	"os"
	"sort"
	"unicode/utf8"
)

var sep1 string
var sep2 string
var joinType string

func init() {
	flag.StringVar(&sep1, "d1", ",", "field delimiter (',' by default)")
	flag.StringVar(&sep2, "d2", ",", "field delimiter (',' by default)")
	flag.StringVar(&sep2, "d", ",", "field delimiter (',' by default)")
	flag.StringVar(&joinType, "type", "inner", "Type of join (inner, left, right, outer or full)")
}

func main() {
	flag.Parse()

	var fileA, fileB *os.File
	var err error
	var colA, colB string
	// Abre el archivo para ser leído
	if len(flag.Args()) == 4 {
		fileA, err = os.Open(flag.Arg(0))
		if err != nil {
			panic(err)
		}
		fileB, err = os.Open(flag.Arg(2))
		if err != nil {
			panic(err)
		}
		colA = flag.Arg(1)
		colB = flag.Arg(3)
	} else if len(flag.Args()) == 3 {
		fileA = os.Stdin
		fileB, err = os.Open(flag.Arg(1))
		if err != nil {
			panic(err)
		}
		colA = flag.Arg(0)
		colB = flag.Arg(2)
	} else {
		panic("Invalid number of arguments")
	}

	// Configura la lectura del CSV
	readerA := bufio.NewReader(fileA)
	rA := csv.NewReader(readerA)
	rA.Comma, _ = utf8.DecodeRuneInString(sep1)

	// Configura la lectura del otro CSV
	readerB := bufio.NewReader(fileB)
	rB := csv.NewReader(readerB)
	rB.Comma, _ = utf8.DecodeRuneInString(sep2)

	// Abrimos el nuevo CSV en la salida estandar
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	w := csv.NewWriter(writer)
	defer w.Flush()

	// Leemos e imprimimos los encabezados:
	colsA, er := rA.Read()
	colsB, er2 := rB.Read()
	if er != nil {
		panic(er)
	}
	if er2 != nil {
		panic(er2)
	}
	w.Write(append(colsA, colsB...))

	// Se leen los archivos completos:
	linesA, err := rA.ReadAll()
	if err != nil {
		panic(err)
	}
	dataA := common.GenTable(colsA, colA, linesA)
	sort.Sort(common.ByValue(dataA))

	linesB, err := rB.ReadAll()
	if err != nil {
		panic(err)
	}
	dataB := common.GenTable(colsB, colB, linesB)
	sort.Sort(common.ByValue(dataB))

	// Imprime la salida del join
	var lineas chan []string
	switch joinType {
	case "inner":
		lineas = innerJoin(dataA, dataB)
	case "left":
		lineas = leftJoin(dataA, dataB)
	case "right":
		lineas = rightJoin(dataA, dataB)
	case "full", "outer":
		lineas = fullJoin(dataA, dataB)
	default:
		panic(fmt.Errorf("'%v' is an invalid join type.", joinType))
	}

	v, ok := <-lineas
	for ok {
		w.Write(v)
		v, ok = <-lineas
	}
}

func innerJoin(dataA, dataB []common.TableData) chan []string {
	sChan := make(chan []string)
	go func() {
		kB := 0
		for _, v := range dataA {
			for kB < len(dataB) && dataB[kB].Value < v.Value {
				kB++
			}
			if kB < len(dataB) && dataB[kB].Value == v.Value {
				sChan <- append(v.Data, dataB[kB].Data...)
			}
		}
		close(sChan)
	}()
	return sChan
}

func leftJoin(dataA, dataB []common.TableData) chan []string {
	sChan := make(chan []string)
	go func() {
		kB := 0
		for _, v := range dataA {
			for kB < len(dataB) && dataB[kB].Value < v.Value {
				kB++
			}
			if kB < len(dataB) && dataB[kB].Value == v.Value {
				sChan <- append(v.Data, dataB[kB].Data...)
			} else {
				sChan <- append(v.Data, make([]string, len(dataB[0].Data))...)
			}
		}
		close(sChan)
	}()
	return sChan
}

func rightJoin(dataA, dataB []common.TableData) chan []string {
	sChan := make(chan []string)
	go func() {
		kA := 0
		for _, v := range dataB {
			for kA < len(dataA) && dataA[kA].Value < v.Value {
				kA++
			}
			if kA < len(dataA) && dataA[kA].Value == v.Value {
				sChan <- append(dataA[kA].Data, v.Data...)
			} else {
				sChan <- append(make([]string, len(dataA[0].Data)), v.Data...)
			}
		}
		close(sChan)
	}()
	return sChan
}

func fullJoin(dataA, dataB []common.TableData) chan []string {
	sChan := make(chan []string)
	go func() {
		kA := 0
		kB := 0
		for kA < len(dataA) || kB < len(dataB) {
			if kA < len(dataA) &&
				(kB >= len(dataB) || dataA[kA].Value < dataB[kB].Value) {
				// "Left join"
				sChan <- append(dataA[kA].Data, make([]string, len(dataB[0].Data))...)
				kA++
			} else if kB < len(dataB) &&
				(kA >= len(dataA) || dataA[kA].Value > dataB[kB].Value) {
				// "Right join"
				sChan <- append(make([]string, len(dataA[0].Data)), dataB[kB].Data...)
				kB++
			} else {
				/* if kA < len(dataA) && kB < len(dataB) && dataA[kA].Value == dataB[kB].Value */
				sChan <- append(dataA[kA].Data, dataB[kB].Data...)
				kA++
				kB++
			}
		}
		close(sChan)
	}()
	return sChan
}
