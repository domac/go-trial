package main

import (
	"fmt"
	"unsafe"
)

func main() {
	v1 := struct {
		a byte
		b byte
		c int32
	}{}

	v2 := struct {
		a byte
		b byte
	}{}

	v3 := struct {
		a byte
		b []int // ptr / len /cap
		c byte
	}{}

	v4 := struct {
		a byte
		c byte
		b []int // ptr / len /cap
	}{}

	v5 := struct {
		a byte
		c int32
		b byte
	}{}

	fmt.Printf("v1: %d , %d \n", unsafe.Alignof(v1), unsafe.Sizeof(v1)) // 4+4
	fmt.Printf("v2: %d , %d \n", unsafe.Alignof(v2), unsafe.Sizeof(v2)) // 1+1
	fmt.Printf("v3: %d , %d \n", unsafe.Alignof(v3), unsafe.Sizeof(v3)) // 24 + 8 + 8
	fmt.Printf("v4: %d , %d \n", unsafe.Alignof(v4), unsafe.Sizeof(v4)) // 24 + 8
	fmt.Printf("v5: %d , %d \n", unsafe.Alignof(v5), unsafe.Sizeof(v5)) // 4 +4+ 4
}
