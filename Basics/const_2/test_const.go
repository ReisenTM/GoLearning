package main

import "fmt"

// const定义枚举类型
const (
	// 可以添加一个关键词iota，开启枚举自增
	// Beijing =0
	Beijing = iota * 10
	ShangHai
	Nanjing
)

// iota只能在const块内使用
const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, k
)

func main() {
	fmt.Printf("Beijing:%d\n", Beijing)
}
