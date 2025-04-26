package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 加奇减偶
type Student struct {
	Name   string
	Age    int
	gender string
}

func stTest(w http.ResponseWriter, r *http.Request) {
	//2.解析模板
	ptmp, err := template.ParseFiles("./templates/struct.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmp := Student{
		Name:   "张三",
		Age:    20,
		gender: "male",
	}
	//3.渲染模板
	err = ptmp.Execute(w, tmp)
	if err != nil {
		fmt.Println("render failed,", err)
		return
	}
}
