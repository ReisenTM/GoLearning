package main

import "fmt"

type Human struct {
	name string
	age  int8
}

func (human *Human) introduce() {
	fmt.Printf("I am %v\n", human)
}

type Student struct {
	Human // 表示继承
	id    string
}

func (stu *Student) introduce() {
	fmt.Println("student from human")
}

func (stu *Student) print() {
	fmt.Println("my name is", stu.name)
	fmt.Println("my age is", stu.age)
	fmt.Println("my id is", stu.id)
}

//func main() {
	human := Human{"ada", 1}
	human.introduce()
	// 直接定义
	stu := Student{Human{"palu", 18}, "19323131"}
	stu.introduce()
	var stu2 Student
	stu2.age = 22
	stu2.name = "lcy"
	stu2.id = "202334070153"
	stu2.introduce()
	stu2.print()
}
