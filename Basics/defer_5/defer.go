package main

func func1() {
	println("fun1")
}

func func2() {
	println("func2")
}

func func3() {
	println("func3")
}

// 函数里，return先执行，defer后执行
func deferTest() {
	defer func1()
	return
}

func main() {
	println("heloo go")
	// defer类似析构函数
	// defer的调用顺序类似压栈
	defer println("main end")
	defer func1()
	defer func2()
	defer func3()
}
