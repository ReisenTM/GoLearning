package main

import (
	"fmt"

	"io"
	"strings"

	"net/http"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func get() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	//记得关闭
	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	//回应内容
	fmt.Println(string(content))
}
func post() {
	body := strings.NewReader("hello world")
	r, err := http.Post("http://httpbin.org/post", "jim", body)
	if err != nil && err != io.EOF {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("回复内容\n", string(content))
}

// 没有封装好的put和delete方法，需要自己仿照post和get实现
func put() {
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
func del() {
	request, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

//func main() {
//	del()
//}
