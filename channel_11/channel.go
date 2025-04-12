package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			temp := x
			x = y
			y = temp + y
		case <-quit:
			fmt.Println("quit")
			return

		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for range 10 {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacii(c, quit)
	// go func() {
	// 	// channel会阻塞的形式实现数据同步
	// 	c <- 666
	// }()
	// num := <-c
	// fmt.Println("num:", num)
	//
	// // 有缓冲channel
	// channel := make(chan int, 3)
	// go func() {
	// 	defer fmt.Println("GoRoutine结束")
	// 	for i := 1; ; i++ {
	// 		channel <- i
	// 		fmt.Println("当前gorotine发送数据", i, "Channel size:", len(channel), "channel capacity:", cap(channel))
	// 	}
	// }()
	//
	// time.Sleep(2 * time.Second)
	// // for range 4 {
	// // 	num := <-channel
	// // 	fmt.Println("当前main接受数据", num, "Channel size:", len(channel), "channel capacity:", cap(channel))
	// // }
	// for num := range channel {
	// 	fmt.Println("当前main接受数据", num, "Channel size:", len(channel), "channel capacity:", cap(channel))
	// }
	// fmt.Println("Main结束")
}
