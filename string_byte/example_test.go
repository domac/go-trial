package main

import (
	"strings"
	"testing"
)

//go test -v -bench . -benchmem

var s = strings.Repeat("a", 1024)

func test() {
	b := []byte(s)
	_ = string(b)
}

func test2() {
	b := stringToBytes(s)
	_ = bytesToString(b)
}

//普通的转换测试
func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

//无额外复制的测试
func BenchmarkTestBlock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test2()
	}
}
