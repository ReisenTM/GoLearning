package main

import (
	"fmt"
	"gorm.io/gorm"
)

func main() {
	mydb.Session(&gorm.Session{
		Logger: mydb.Logger,
	})
	var studentList []Student
	mydb.Where("name like ?", "张%").Find(&studentList)

	fmt.Println(studentList)
}
