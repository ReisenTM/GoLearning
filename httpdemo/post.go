package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

func postForm() {
	//http.Post()
	data := make(url.Values)
	data.Add("name", "John Doe")
	data.Add("email", "johndoe@gmail.com")

	u := "http://httpbin.org/post"
	//相当于http.Post(u,content-type,data)里让content-type="application/x-www-form-urlencoded"的封装
	r, _ := http.PostForm(u, data)
	defer r.Body.Close()
	content, _ := io.ReadAll(r.Body)

	fmt.Println(string(content))
}
func postJson() {
	//http.Post()
	sourceData := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{"John Doe", "johndoe@gmail.com"}
	data, _ := json.Marshal(sourceData)
	u := "http://httpbin.org/post"
	r, _ := http.Post(u, "application/json", bytes.NewReader(data))
	defer r.Body.Close()
	content, _ := io.ReadAll(r.Body)

	fmt.Println(string(content))
}
func printReq() {
	// 构造请求
	data := make(url.Values)
	data.Add("name", "John Doe")
	data.Add("email", "johndoe@gmail.com")

	req, _ := http.NewRequest("POST", "http://httpbin.org/post", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// 打印请求
	dump, _ := httputil.DumpRequest(req, true)
	fmt.Println(string(dump))
}
func postFile() {
	//数据缓冲区
	body := &bytes.Buffer{}
	//创建拼接类
	writer := multipart.NewWriter(body)
	//增加字段
	_ = writer.WriteField("name", "John Doe")
	_ = writer.WriteField("email", "johndoe@gmail.com")
	//上传文件 表单名 + 文件名
	uploadFile_1, _ := writer.CreateFormFile("uploadfile_1", "test.jpg")
	file, _ := os.Open("./request.go")
	defer file.Close()
	_, err := io.Copy(uploadFile_1, file)
	if err != nil {
		return
	}
	//关闭拼接
	_ = writer.Close()
	fmt.Println(writer.FormDataContentType())
	fmt.Println(body)
	r, _ := http.Post("http://httpbin.org/post", writer.FormDataContentType(), body)
	defer r.Body.Close()
	content, _ := io.ReadAll(r.Body)
	fmt.Println(string(content))

}

//func main() {
//	postFile()
//}
