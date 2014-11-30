package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var acum float64
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		num, err := strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			panic(err)
		}
		acum += num
	}
	if sc.Err() != nil {
		panic(sc.Err())
	}
	fmt.Println(acum)
}
