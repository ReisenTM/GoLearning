package main

import "fmt"

// interface{} 是万能数据类型，可以看作没有任何方法的接口，所有数据类型都实现了它(Objcet?)
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)
	// interface{} 如何区分此时传入的数据类型是什么？
	// 提供有 类型断言机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg type is string")
		fmt.Printf("arg type :%T\n", value)
	}
}

func main() {
	myFunc("dada")
}
