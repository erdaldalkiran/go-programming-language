package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type dirByte struct {
	path string
	size int64
}

type dirByteLevel struct {
	dirByte
	level int
}

func main() {

	result := make(map[int][]dirByte)

	dirs := os.Args[1:]
	if len(dirs) == 0 {
		dirs = append(dirs, ".")
	}

	var dataChs []chan dirByteLevel

	for _, dir := range dirs {
		ch := make(chan dirByteLevel)
		dataChs = append(dataChs, ch)
		go walkDir(dir, 0, ch)
	}

	out := merge(dataChs...)
	for d := range out {
		result[d.level] = append(result[d.level], d.dirByte)
	}

	for level, dirBytes := range result {
		for _, db := range dirBytes {

			if level < 2 {
				fmt.Printf("level: %d, size of %v %d B\n", level, db.path, db.size)
				fmt.Printf("level: %d, size of %v %.2f GB\n", level, db.path, float64(db.size)/1e9)
			}
		}
	}
}

func walkDir(dir string, level int, out chan dirByteLevel) {
	var result int64 = 0

	var inners []chan dirByteLevel

	for _, entry := range dirents(dir) {

		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			ch := make(chan dirByteLevel)
			inners = append(inners, ch)
			go walkDir(subdir, level+1, ch)
		}

		result += entry.Size()
	}

	innerAll := merge(inners...)

	for dbl := range innerAll {
		if dbl.level == level+1 {
			result += dbl.size
		}
		out <- dbl
	}

	out <- dirByteLevel{dirByte{dir, result}, level}

	close(out)
}

func dirents(dir string) []os.FileInfo {
	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entries, err := f.Readdir(0) // 0 => no limit; read all entries
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		// Don't return: Readdir may return partial results.
	}
	return entries
}

func merge(cs ...chan dirByteLevel) chan dirByteLevel {
	out := make(chan dirByteLevel)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c chan dirByteLevel) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
