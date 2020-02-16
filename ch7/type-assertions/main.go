package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	fmt.Println(w)
	w = os.Stdout
	fmt.Println("value is", w)
	fmt.Printf("type is: %T\n", w)
	f := w.(*os.File)
	fmt.Printf("f: %T\n", f)

	c, ok := w.(*bytes.Buffer)
	fmt.Printf("c: %T ok: %v\n", c, ok)

	var w2 io.Writer
	w2 = os.Stdout
	rw := w2.(io.ReadWriter)
	fmt.Printf("w2: %T\n", w2)
	fmt.Printf("rw: %T\n", rw)
}
