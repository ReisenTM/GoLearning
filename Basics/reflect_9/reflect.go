package main

import (
	"fmt"
	"reflect"
	"student/student"
)

type User struct {
	Name string
	age  int
}

// 一个变量内部包含类型和value两个属性，反射就是运行时动态检查变量的类型和值
func GetType(arg any) {
	fmt.Println("type of arg is:", reflect.TypeOf(arg).Name())
	fmt.Println("value of arg is:", reflect.ValueOf(arg))
}

func GetStructInfo(st any) {
	rf := reflect.TypeOf(st).Elem()
	for i := 0; i < rf.NumField(); i++ {
		field := rf.Field(i)
		fmt.Println("Name of field:", field.Name)
		fmt.Println("Type of field:", field.Type)
		// exported :是否在其他包可见
		fmt.Println("is Export", field.IsExported())
		fmt.Println("-----")
	}
}

/*
	func main() {
		user := User{"zhangsan", 22}
		GetType(user)
		fmt.Println("-------------")
		GetStructInfo(&user)
	}
*/
func main() {
	// 这种叫做 unkeyed literal，Go 现在鼓励显式写 key
	stu_1 := student.Student{"帕鲁", "雄性"}
	student.FindTag(stu_1)
	stu := student.Student{Name: "帕鲁", Gender: "雄性"}
	student.FindTag(stu)

	movie := student.Movie{
		Name:   "命运之夜：天之杯",
		Time:   2016,
		Actors: []string{"saber", "archer", "lancer", "caster", "assassins", "basaker"},
	}
	JsonBytes := student.MovieInfo(movie)
	c_movie := student.MovieCollect(JsonBytes)
	fmt.Printf("json to struct result:%v\n", c_movie)
}
