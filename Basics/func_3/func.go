package main

import "fmt"

func func_1(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	return 10
}

func func_2(a string, b int) (int, int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	return 10, 20
}

func func_2_arrange(a string, b int) (r1 int, r2 int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	// 可以直接在函数内给有名称的返回值赋值
	r1 = 10
	r2 = 10
	return 10, 20
}

func func_2_arrange2(a string, b int) (r1, r2 int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	return 10, 20
}

func main() {
	fmt.Println("Func-1--------")
	f1_res := func_1("test1", 333)
	fmt.Println("f1 result:", f1_res)
	fmt.Println("Func-2--------")
	f2_res1, f2_res2 := func_2("test2", 666)
	fmt.Println("f1 result:", f2_res1, "f1 res2:", f2_res2)
	// 匿名函数
	a := func() {
		defer fmt.Println("A.defer")

		fmt.Println("A")
	}
	a()

	// 匿名函数直接调用
	func() {
		defer fmt.Println("B.defer")

		fmt.Println("B")
	}() // 多一个"()"
}
