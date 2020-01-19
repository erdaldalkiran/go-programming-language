package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	elapsed := time.Since(start).Seconds()
	fmt.Printf("%.2fs elaspsed since start \n", elapsed)
}

func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("error while getting url %s: %v", url, err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	defer resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("error while reading %s: %v", url, err)
		return
	}
	elapsed := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%d bytes read from %s in %.2f seconds", nbytes, url, elapsed)
}
