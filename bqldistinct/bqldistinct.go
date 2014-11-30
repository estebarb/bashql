package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	m := make(map[string]bool)
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		m[sc.Text()] = true
	}
	if sc.Err() != nil {
		panic(sc.Err())
	}
	fmt.Println(len(m))
}
