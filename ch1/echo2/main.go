package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string

	for i, arg := range os.Args[1:] {
		s += sep + strconv.Itoa(i) + ":" + arg
		sep = "   "
	}

	fmt.Println(s)
}
