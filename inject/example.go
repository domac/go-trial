package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type Demo struct {
}

type TestObj interface{}

func (d *Demo) Say(name string, version float64, obj TestObj) {
	fmt.Printf("Coding with %s, obj is %v, version is %f!\n", name, obj, version)
}

func main() {
	testDemo := new(Demo)
	inj := inject.New()
	inj.Map("Golang")
	inj.MapTo("语言", (*TestObj)(nil))
	inj.Map(10.0)
	f := testDemo.Say
	inj.Invoke(f)
}
