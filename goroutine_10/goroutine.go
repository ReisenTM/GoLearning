package main

import (
	"fmt"
	"runtime"
	"time"
)

// go天生支持并发
func test() {
	defer println("test routine exited")

	i := 0
	for {
		i++
		if i == 10 {
			// 退出协程
			runtime.Goexit()
		}
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// go test()
	// i := 0
	// for {
	// 	i++
	// 	fmt.Println("main:", i)
	// 	time.Sleep(1 * time.Second)
	// }
}
