package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

//虽然反射提供的API远多于我们讲到的，我们前面的例子主要是给出了一个方向，通过反射可以实现哪些功能。
//反射是一个强大并富有表达力的工具，但是它应该被小心地使用，原因有三。

//第一个原因是，基于反射的代码是比较脆弱的。对于每一个会导致编译器报告类型错误的问题，在反射中都有与之相对应的问题，
//不同的是编译器会在构建时马上报告错误，而反射则是在真正运行到的时候才会抛出panic异常，可能是写完代码很久之后的时候了，
//而且程序也可能运行了很长的时间。
//避免这种因反射而导致的脆弱性的问题的最好方法是将所有的反射相关的使用控制在包的内部，
//如果可能的话避免在包的API中直接暴露reflect.Value类型，这样可以限制一些非法输入。
//如果无法做到这一点，在每个有风险的操作前指向额外的类型检查。以标准库中的代码为例，
//当fmt.Printf收到一个非法的操作数是，它并不会抛出panic异常，而是打印相关的错误信息。
//程序虽然还有BUG，但是会更加容易诊断。

//第二个原因是，即使对应类型提供了相同文档，但是反射的操作不能做静态类型检查，而且大量反射的代码通常难以理解。
//总是需要小心翼翼地为每个导出的类型和其它接受interface{}或reflect.Value类型参数的函数维护说明文档。

//第三个原因，基于反射的代码通常比正常的代码运行速度慢一到两个数量级。
//对于一个典型的项目，大部分函数的性能和程序的整体性能关系不大，所以使用反射可能会使程序更加清晰。
//测试是一个特别适合使用反射的场景，因为每个测试的数据集都很小。但是对于性能关键路径的函数，最好避免使用反射。

// Any formats any value as a string.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 10)
	default:
		return v.Type().String() + " value"
	}
}

func main() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))                  // "1"
	fmt.Println(Any(d))                  // "1"
	fmt.Println(Any([]int64{x}))         // "[]int64 0x842350927976"
	fmt.Println(Any([]time.Duration{d})) // "[]time.Duration 0x842350928032"
}
