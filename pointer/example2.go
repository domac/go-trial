package main

import (
	"fmt"
	"unsafe"
)

//返回局部变量指针是安全的,编译器会根据需要将其分配在 GC Heap 上。

//将 Pointer 转换成 uintptr,可变相实现指针运算。
func main() {
	d := struct {
		s string
		x int
	}{"abc", 100}

	p := uintptr(unsafe.Pointer(&d))
	p += unsafe.Offsetof(d.x)

	p2 := unsafe.Pointer(p)
	px := (*int)(p2)
	*px = 200

	fmt.Printf("%#v\n", d)
}
