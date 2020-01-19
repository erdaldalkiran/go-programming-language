package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			defer f.Close()
			if err != nil {
				fmt.Printf("error openingcd file:%s error:%v\n", file, err)
				continue
			}
			countLines(f, counts)
		}
	}

	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%v:\t%d\n", line, count)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}
