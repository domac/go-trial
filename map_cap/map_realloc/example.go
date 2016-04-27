package main

import (
	"runtime/debug"
	"time"
)

const cap = 1000000

var dict = make(map[int][100]byte, cap)

func test() {
	for i := 0; i < cap; i++ {
		dict[i] = [100]byte{}
	}

	for k := range dict {
		delete(dict, k)
	}

	dict = nil
}

func main() {
	test()

	for i := 0; i < 20; i++ {
		debug.FreeOSMemory()
		time.Sleep(time.Second)
	}
}
