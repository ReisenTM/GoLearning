package main

import "fmt"

// 动态数组，传递是切片副本,使用起来像指针
// go只有值传递
func printArray(myArr []int) {
	//'_'表示匿名的变量
	for _, value := range myArr {
		fmt.Println("value:", value)
	}
}

func main() {
	myArray1 := []int{1, 2, 3, 4, 10} // 动态数组 ，切片slice
	fmt.Printf("myArray1 type is:%T\n", myArray1)
	printArray(myArray1)
}
