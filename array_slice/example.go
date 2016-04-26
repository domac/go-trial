package main

import (
	"fmt"
)

const cap = 1024

func array() [cap]int {
	var d [cap]int
	for i := 0; i < len(d); i++ {
		d[i] = 1
	}
	return d
}

func slice() []int {
	d := make([]int, cap)
	for i := 0; i < len(d); i++ {
		d[i] = 1
	}
	return d
}

func main() {
	fmt.Println(array())
	fmt.Println(slice())
}
