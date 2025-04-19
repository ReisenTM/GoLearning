package main

import (
	"init_4/lib1"
	// 可以给包起别名
	mylib "init_4/lib2"
	//匿名导包，前面加'_',可以只调用包的init函数
	//_"package/path"
	//导入包内全部内容 (相当于python里的from xxx import *)
	//不推荐，同名函数会导致冲突
	//."package/path"
)

/*
*	go语言中，包内函数名称首字母的大小写决定了它的公开性
* 首字母小写，如:privatFunc 表示只在包内使用的函数  首字母大写，如PublicFunc 可以被包外的程序调用
* */
func main() {
	lib1.Lib1_test()
	mylib.Lib2_test()
}
