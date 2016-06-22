package main

import (
	"fmt"
)

type Demo struct {
}

type TestObj interface{}

func (d *Demo) Say(name string, version float64, obj TestObj) {
	fmt.Printf("Coding with %s, obj is %v, version is %f!\n", name, obj, version)
}

func main() {
	testDemo := new(Demo)
	testDemo.Say("Golang", 10.0, "语言")
}
