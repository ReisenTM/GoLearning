package main

import "fmt"

// 数组直接传参不同长度算不同的类型,而且是值传递
func printArray(myArray [4]int) { // X myArray1
	for index, value := range myArray {
		fmt.Println("index=", index, "value", value)
	}
}

func main() {
	var myArray1 [10]int
	myArray := [10]int{1, 3, 5, 7}
	// 1. for 遍历
	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray1[i])
	}
	// 2. range遍历
	for index, value := range myArray {
		fmt.Println("index=", index, "value", value)
	}

	// 查看数组的数据类型
	fmt.Printf("myArray1 types=%T\n", myArray1)
	fmt.Printf("myArray types=%T\n", myArray)
}
