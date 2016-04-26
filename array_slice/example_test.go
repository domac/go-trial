package main

import (
	"testing"
)

//go test -v -bench . -benchmem

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slice()
	}
}

//array 非但拥有更好的性能，还避免了堆内存分配，也就是说减轻了 GC 压力
//函数 array 返回值的复制只需用 "CX + REP" 指令就可完成
//整个 array 函数完全在栈上完成，而 slice 函数则需执行 makeslice，继而在堆上分配内存，这就是问题所在
//对于一些短小的对象，复制成本远小于在堆上分配和回收操作
//Go Proverbs: A little copying is better than a little dependency.
