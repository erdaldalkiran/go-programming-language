package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var counter int

func main() {
	http.HandleFunc("/count", count)
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("localhost:8002", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counter++
	mu.Unlock()

	fmt.Println("handle")
	fmt.Fprintf(w, "%q received", r.URL.Path)
}

func count(w http.ResponseWriter, r *http.Request) {
	fmt.Println("count")
	fmt.Fprintf(w, "all request count %d", counter)
}
