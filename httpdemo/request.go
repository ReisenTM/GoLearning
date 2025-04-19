package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func printBody(req *http.Response) {
	defer req.Body.Close()
	content, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

// 带查询参数请求
func requestByParams() {
	req, err := http.NewRequest("GET", "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	//设置url参数
	params := make(url.Values)
	params.Set("param1", "value1")
	params.Set("param2", "value2")
	params.Add("param1", "value3")
	fmt.Println(req.Method, params)
	fmt.Println(params.Encode())
	//编码成param1=value1&param1=value3格式
	req.URL.RawQuery = params.Encode()
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	printBody(r)
}

// 定制请求头
func requestByHeaders() {
	req, err := http.NewRequest("GET", "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	//可以通过修改User-Agent简单绕过反爬
	req.Header.Set("User-Agent", "Safari")
	req.Header.Add("header1", "value1")
	req.Header.Add("header2", "value2")
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	printBody(r)
}

//func main() {
//	requestByParams()
//}
