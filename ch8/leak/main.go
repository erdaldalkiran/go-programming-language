package main

import "fmt"

func main() {
	done := make(chan struct{})

	result := func() string {
		res := make(chan string) // goroutine leak
		//res := make(chan string, 3)

		go func() {
			res <- "1"
		}()
		go func() {
			res <- "2"
		}()
		go func() {
			res <- "3"
		}()

		return <-res
	}()

	fmt.Println(result)
	go func() {
		for i := 0; i < 1000; i++ {

		}

		done <- struct{}{}
	}()
	<-done
}
