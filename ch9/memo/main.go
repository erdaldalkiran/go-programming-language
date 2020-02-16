package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex //guards cache
	cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	defer memo.mu.Unlock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
		fmt.Println("###", key)
	}

	return res.value, res.err
}

func main() {
	incomingURLs := []string{"https://golang.org", "https://godoc.org", "https://play.golang.org", "http://gopl.io",
		"https://golang.org", "https://godoc.org", "https://play.golang.org", "http://gopl.io"}

	m := New(httpGetBody)
	var n sync.WaitGroup
	for _, url := range incomingURLs {
		n.Add(1)

		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}

			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))

			n.Done()
		}(url)
	}

	n.Wait()
}
