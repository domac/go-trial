package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	//因为 reflect.TypeOf 返回的是一个动态类型的接口值, 它总是返回具体的类型.
	//因此, 下面的代码将打印 "*os.File" 而不是 "io.Writer".
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
}
