package main

import "fmt"

func main() {
	// 列表初始化
	slice_1 := []int{1, 3, 5, 4}
	// 声明但不分配空间
	var slice_2 []int
	fmt.Printf("slice2:%#v\n", slice_2)
	// 使用make开辟空间
	slice_3 := make([]int, 3)
	fmt.Printf("slice3:%#v\n", slice_3)
	fmt.Printf("len=%d,slice=%v\n", len(slice_1), slice_1)
}
