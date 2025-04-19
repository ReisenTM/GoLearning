package main

import (
	"fmt"
	"gorm.io/gorm"
)

type Student struct {
	//利用标签可以设置gorm键值属性
	Name   string `gorm:"size:16;comment:姓名"`
	Age    int    `gorm:"size:3;comment:年龄"`
	Id     uint   `gorm:"size:10;comment:ID"` //默认为主键
	Gender bool   `gorm:"size:2"`
	//指针的作用是可以设置属性为空值
	Email *string `gorm:"size:128;comment:邮箱"`
}

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Printf("hook function before create\n")
	eml := "s213131@qq.com"
	s.Email = &eml
	return nil
}

func update() {
	//单条记录的查询
	mydb = mydb.Session(&gorm.Session{
		Logger: mydbLogger,
	})
	/*var student Student
	mydb.Take(&student)
	fmt.Println(student)
	//数据不干净可能导致查询出错
	student = Student{}
	//?是占位符
	mydb.Take(&student, "name=?", "张3")
	//防止sql注入，切忌自己拼接字符串
	//sql注入原理就是字符串转义
	fmt.Println(student)
	*/
	//查询多条记录
	var studentList []Student
	count := mydb.Find(&studentList).RowsAffected
	for _, student := range studentList {
		fmt.Println(student)
	}
	fmt.Printf("一共有%d条记录:", count)
	//单个记录所有字段的更新
	var saveStudent Student
	saveStudent.Name = "帕鲁"
	saveStudent.Id = 1
	saveStudent.Age = 100
	mydb.Save(&saveStudent)

	newEmail := "30343242@gmail.com"
	//批量更新
	//单个字段更新
	mydb.Find(&studentList, "age=?", 10).Select("email").Updates(map[string]interface{}{
		"name":  "hello",
		"age":   18,
		"email": newEmail})

	mydb.Find(&studentList, "age=?", 10).Update("name", "批量更新测试2")
	//注意 使用 struct 更新时, GORM 将只更新非零值字段。 可以用 map 来更新属性，或者使用 Select 声明字段来更新
	var student Student
	mydb.Find(&studentList).Delete(&student)

	mydb.Delete(&student, []int{1, 2, 3})
	// DELETE FROM users WHERE id IN (1,2,3);

}

func create() {
	//记录的插入
	//mydb.AutoMigrate(&Student{})
	email := "30872223@qq.com"
	/*err := mydb.Create(&Student{
		Name:   "张三",
		Age:    18,
		Gender: true,
		Email:  &email,
	}).Error
	if err != nil {
		fmt.Println(err)
	}*/
	mydb = mydb.Session(&gorm.Session{
		Logger: mydbLogger,
	})
	var stuList []Student
	for i := 0; i < 10; i++ {
		stuList = append(stuList, Student{
			Name:   fmt.Sprintf("张%d", i+11),
			Age:    i + 1,
			Gender: true,
			Email:  &email,
		})
	}
	err := mydb.Create(&stuList).Error
	if err != nil {
		fmt.Println(err)
	}

}
