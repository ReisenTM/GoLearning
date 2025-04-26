package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	//2.解析模板
	ptmp, err := template.ParseFiles("./templates/hello.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	//3.渲染模板
	//字符串渲染
	err = ptmp.Execute(w, "测试测试")
	if err != nil {
		fmt.Println("render failed,", err)
		return
	}
}
