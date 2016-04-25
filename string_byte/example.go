package main

import (
	"fmt"
	"strings"
	"unsafe"
)

//无额外复制的转换
func stringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	s := strings.Repeat("abc", 3)
	b := stringToBytes(s)
	s2 := bytesToString(b)

	fmt.Println(b, s2)
}
