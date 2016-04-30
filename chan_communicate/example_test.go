package main

import (
	"testing"
)

func BenchmarkChanCounter(b *testing.B) {
	c := chanCounter()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = <-c
	}
}

func BenchmarkMutexCounter(b *testing.B) {
	f := mutexCounter()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = f()
	}
}

func BenchmarkAtomicCounter(b *testing.B) {
	f := atomicCounter()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = f()
	}
}
