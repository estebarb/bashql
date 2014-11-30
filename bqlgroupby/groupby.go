package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/estebarb/bashql/common"
	"io"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type stringslice []string

// Now, for our new type, implement the two methods of
// the flag.Value interface...
// The first method is String() string
func (ss *stringslice) String() string {
	return fmt.Sprintf("%v", *ss)
}

// The second method is Set(value string) error
func (ss *stringslice) Set(value string) error {
	//fmt.Printf("%s\n", value)
	*ss = append(*ss, value)
	return nil
}

var groupBys stringslice
var reduceBys stringslice
var reducers stringslice

func init() {
	flag.Var(&groupBys, "g", "List of columns to be grouped")
	flag.Var(&reduceBys, "c", "List of columns to be reduced")
	flag.Var(&reducers, "f", "Reduction to be applied to column")
}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		panic("Not enought arguments")
	}

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
	w.Write(append(groupBys, reduceBys...))

	// cargamos todos los datos
	rawData, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	data := common.GenMultiValueTable(columnas, groupBys, rawData)
	sort.Sort(common.ByMultiValue(data))

	agrupando := false
	currentReducers := make([]*exec.Cmd, len(reducers))
	outputBuffers := make([]bytes.Buffer, len(reducers))
	inputPipe := make([]io.WriteCloser, len(reducers))
	for i := 0; i < len(data); i++ {
		if agrupando {
			// Alimenta los reductores
			row := data[i]
			for k, v := range inputPipe {
				fmt.Fprintln(v, common.GetValue(reduceBys[k], row.Data, columnas))
			}
		} else {
			// Crea nuevos reductores
			currentReducers = make([]*exec.Cmd, len(reducers))
			outputBuffers = make([]bytes.Buffer, len(reducers))
			inputPipe = make([]io.WriteCloser, len(reducers))
			for k, _ := range currentReducers {
				currentReducers[k] = exec.Command(reducers[k])
				currentReducers[k].Stdout = &outputBuffers[k]
				inputPipe[k], err = currentReducers[k].StdinPipe()
				defer inputPipe[k].Close()
				if err != nil {
					panic(err)
				}
				err := currentReducers[k].Start()
				if err != nil {
					panic(err)
				}
			}
			agrupando = true
			i--
			continue
		}
		// Si la siguiente fila corresponde a otro grupo entonces
		// hay que recuperar los resultados de la fila actual
		if i+1 == len(data) || !(common.ByMultiValue(data)).Equal(i, i+1) {
			agrupando = false
			valsReducidos := make([]string, len(reduceBys))
			for k, v := range currentReducers {
				inputPipe[k].Close()
				err := v.Wait()
				if !v.ProcessState.Success() {
					panic(err)
				}
				valsReducidos[k] = strings.TrimSpace(outputBuffers[k].String())
			}
			w.Write(append(data[i].MultiValue, valsReducidos...))
		}
	}
}
