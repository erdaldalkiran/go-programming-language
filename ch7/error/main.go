package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

var ErrNotExist = errors.New("file does not exist")

// IsNotExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist. It is satisfied by
// ErrNotExist as well as some syscall errors.
func IsNotExist(err error) bool {
	if pe, ok := err.(*os.PathError); ok {
		err = pe.Err
	}
	return err == syscall.ENOENT || err == ErrNotExist
}

func main() {
	var c interface{}
	c, _ = os.Open("main.go")
	switch x := c.(type) {
	default:
		fmt.Printf("%T %T\n", x, c)
	}
	fmt.Printf("%T\n", c)
	_, err := os.Open("/main.go")
	fmt.Printf("%#v\n", err)
	fmt.Printf("%#v\n", err.(*os.PathError).Err)
	fmt.Println(os.IsNotExist(err)) // "true
}
