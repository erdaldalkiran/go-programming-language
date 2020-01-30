package main

import (
	"flag"
	"fmt"
)

type hede struct {
	value float64
	kind  string
}

type hedeFlag struct {
	value hede
}

func (h *hedeFlag) Set(s string) error {
	var value float64
	var kind string

	_, err := fmt.Sscanf(s, "%f%s", &value, &kind)
	if err != nil {
		return err
	}

	h.value = hede{
		value: value,
		kind:  kind,
	}

	return nil
}

func (h *hedeFlag) String() string {
	return fmt.Sprintf("%f%s", h.value.value, h.value.kind)
}

func HedeFlag(name string, value hede, usage string) *hede {
	h := hedeFlag{value}
	flag.CommandLine.Var(&h, name, usage)
	return &h.value
}

var hFlag = HedeFlag("h", hede{10, "ciko"}, "enter some string and see what happens")

func main() {
	flag.Parse()
	fmt.Println(*hFlag)
	fmt.Println(hFlag)
}
