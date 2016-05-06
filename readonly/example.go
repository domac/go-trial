package main

import (
	"errors"
	"fmt"
	"strconv"
)

//检查只读变量的流失

func main() {
	_, err := strconv.ParseInt("xxx", 10, 64)
	fmt.Println(err.(*strconv.NumError).Err == strconv.ErrSyntax)

	strconv.ErrSyntax = errors.New("test error")

	switch e := err.(*strconv.NumError); e.Err {
	case strconv.ErrSyntax:
		fmt.Println(e)
	default:
		fmt.Println("unknow")
	}
}
