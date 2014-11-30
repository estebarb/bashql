package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var acum float64
	acum = math.Inf(1)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		num, err := strconv.ParseFloat(sc.Text(), 64)
		if err != nil {
			panic(err)
		}
		acum = math.Min(acum, num)
	}
	if sc.Err() != nil {
		panic(sc.Err())
	}
	fmt.Println(acum)
}
