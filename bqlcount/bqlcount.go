package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	acum := 0
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		sc.Text()
		acum++
	}
	if sc.Err() != nil {
		panic(sc.Err())
	}
	fmt.Println(acum)
}
