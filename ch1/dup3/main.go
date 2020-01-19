package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Printf("error reading file:%s error: %v \n", file, err)
			continue
		}

		for _, line := range strings.Split(string(data), "\r\n") {
			counts[line]++
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%v:\t%d\n", line, count)
		}
	}
}
