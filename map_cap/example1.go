package main

import (
	"runtime"
	"time"
)

const cap = 500000

var d interface{}

func value() interface{} {
	m := make(map[int]int, cap)
	for i := 0; i < cap; i++ {
		m[i] = i
	}
	return m
}

func pointer() interface{} {
	m := make(map[int]*int, cap)
	for i := 0; i < cap; i++ {
		v := i
		m[i] = &v
	}
	return m
}

// go build -o test && GODEBUG="gctrace=1" ./test

// 两次输出里 GC 所占时间百分比，就可看出 “巨大” 差异

func main() {

	//d = pointer()
	d = value()

	for i := 0; i < 20; i++ {
		runtime.GC()
		time.Sleep(time.Second)
	}
}
