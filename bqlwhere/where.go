package main

import (
	"bufio"
	"encoding/csv"
	"github.com/estebarb/bashql/common"
	"os"
	"fmt"
	"regexp"
	"strconv"
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
	if len(argumentos)%3 != 0{
		panic(fmt.Errorf("Must be called %v columnA operator argument.", os.Args[0]))
	}

	// Abrimos el nuevo CSV en la salida estandar
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	w := csv.NewWriter(writer)
	defer w.Flush()

	// Imprimimos las cabeceras:
	w.Write(columnas)

	col, er := r.Read()
	for er == nil {
		relevante := true
		
		// Determina si la fila es relevante
		for i := 0; i < len(argumentos); i += 3{
			valorA := common.GetValue(argumentos[i], col, columnas)
			valorB := argumentos[i+2]
			numA, errA := strconv.ParseFloat(valorA, 64)
			numB, errB := strconv.ParseFloat(valorB, 64)
			
			operador := argumentos[i+1]
			if operador[0] == 'c'{
				valorB = common.GetValue(valorB, col, columnas)
				numB, errB = strconv.ParseFloat(valorB, 64)
				operador = operador[1:]
			}
			
			switch(operador){
				case "=":
					if errA != nil || errB != nil{
						// Comparación de strings
						relevante = relevante && (valorA == valorB)
					} else {
						// Comparación de números
						relevante = relevante && (numA == numB)
					}
				case "<":
					if errA != nil || errB != nil{
						// Comparación de strings
						panic(fmt.Errorf("'<' not available for strings.", ))
					} else {
						// Comparación de números
						relevante = relevante && (numA < numB)
					}
				case ">":
					if errA != nil || errB != nil{
						// Comparación de strings
						panic(fmt.Errorf("'>' not available for strings.", ))
					} else {
						// Comparación de números
						relevante = relevante && (numA > numB)
					}
				case "<=":
					if errA != nil || errB != nil{
						// Comparación de strings
						panic(fmt.Errorf("'<=' not available for strings.", ))
					} else {
						// Comparación de números
						relevante = relevante && (numA <= numB)
					}
				case ">=":
					if errA != nil || errB != nil{
						// Comparación de strings
						panic(fmt.Errorf("'>=' not available for strings.", ))
					} else {
						// Comparación de números
						relevante = relevante && (numA >= numB)
					}
				case "!=":
					if errA != nil || errB != nil{
						// Comparación de strings
						relevante = relevante && (valorA != valorB)
					} else {
						// Comparación de números
						relevante = relevante && (numA != numB)
					}
				case "like":
					expreg := regexp.MustCompile(valorB)
					relevante = relevante && expreg.MatchString(valorA)
				case "unlike":
					expreg := regexp.MustCompile(valorB)
					relevante = relevante && !expreg.MatchString(valorA)
				default:
					panic(fmt.Errorf("'%v' not recognized as a valid operator.", ))
			}
		}
		
		// Si la fila es relevante entonces se imprime
		if relevante {
			w.Write(col)
		}
		col, er = r.Read()
	}
}
