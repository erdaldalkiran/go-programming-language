package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		var x, y int
		go func() {
			x = 1                   // A1
			fmt.Print("y:", y, " ") // A2
		}()
		go func() {
			y = 1                   // B1
			fmt.Print("x:", x, " ") // B2
		}()

		fmt.Println()
		time.Sleep(1 * time.Millisecond)
	}

}
