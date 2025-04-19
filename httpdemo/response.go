package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io"
	"net/http"
)

func responseBody(r *http.Response) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(content))
}

// 状态
func getStat(r *http.Response) {
	fmt.Println(r.StatusCode) //状态码
	fmt.Println(r.Status)     //状态信息

}
func header(r *http.Response) {
	//Header是一个map，通过get的方式比通过at的方式好处:大小写忽略
	fmt.Println(r.Header.Get("content-type"))
}

// 获取编码信息
func encoding(r *http.Response) {
	//可以通过网页的头部猜测网页的编码信息
	//引入包:golang.org/x/net/html
	bufReader := bufio.NewReader(r.Body)
	bytes, _ := bufReader.Peek(1024) //预读,不会移动reader的读取位置
	//提供content-type的原因：Content-Type 提供的上下文（如语言区域）能显著提高准确性，检测优先级高
	res, _, _ := charset.DetermineEncoding(bytes, r.Header.Get("content-type"))
	fmt.Println(res)
	//不解码结果
	//rawBody, _ := io.ReadAll(bufReader)
	//fmt.Println(string(rawBody))
	//fmt.Println("------------------")
	//引入包:golang.org/x/text/transform，解码
	bodyReader := transform.NewReader(bufReader, res.NewDecoder())
	content, _ := io.ReadAll(bodyReader)
	fmt.Println(string(content))
}

//func main() {
//	r, err := http.Get("https://google.com")
//	if err != nil {
//		panic(err)
//	}
//	defer func() { _ = r.Body.Close() }()
//	encoding(r)
//}
