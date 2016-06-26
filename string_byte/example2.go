package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := []int{1, 2, 3, 4, 5}
	p := &x[0]

	ip := unsafe.Pointer(uintptr(unsafe.Pointer(p)) + unsafe.Sizeof(x[0]))
	p = (*int)(ip)
	fmt.Printf("%d\n", *p)
}
