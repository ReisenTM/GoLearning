>golang.net/http包学习记录
---
请求展示网站，可以将请求的内容以json的形式返回
http://httpbin.org/
**如何打印请求内容？**
```go
// 打印请求  
dump, _ := httputil.DumpRequest(req, true)  
fmt.Println(string(dump))
```
---
## Example
> 流程：发送请求——>接收返回内容
```go
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
```
> net/http默认只封装了get和post方法，put和delete需要自己根据实现
> 实现方法: 
> 	1. NewRequest
> 	2. 设置Content-type
> 	3. 选择Client执行Do方法发送请求
> 	4.接收返回内容
```go
// put sends an HTTP PUT request to http://httpbin.org/put and prints the response.
// The function demonstrates how to make a PUT request with custom headers.
// It sets the Content-Type header but sends no body in this example.
func put() {
	// Create a new PUT request with nil body
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err) // Handle error if request creation fails
	}
	
	// Set the Content-Type header (though we're not sending any content in this example)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Send the request using the default HTTP client
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err) // Handle error if request fails
	}
	defer r.Body.Close() // Ensure the response body is closed when we're done
	
	// Read the entire response body
	content, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err) // Handle error if reading response fails
	}
	
	// Print the response content as string
	fmt.Println(string(content))
}
```
---
## Request
1. 带查询参数请求
```go
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
    //编码成param1=value1&param1=value3格式 
    fmt.Println(params.Encode())  
    req.URL.RawQuery = params.Encode()  
    r, err := http.DefaultClient.Do(req)  
    if err != nil {  
       panic(err)  
    }  
    printBody(r)  
}
//请求内容 ： param1=value1&param1=value3&param2=value2

```
2. 定制请求头
```go
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
```
---
## Response
状态  
```go
func getStat(r *http.Response) {  
    fmt.Println(r.StatusCode) //状态码  200
    fmt.Println(r.Status)     //状态信息  ok 200
}
```
头
```go
func header(r *http.Response) {  
    //Header是一个map，通过get的方式比通过at的方式好处:大小写忽略  
    fmt.Println(r.Header.Get("content-type"))  
}
```
编码
 **需要解码的主要场景**
1. 响应内容使用非UTF-8编码时
    - 网页常用编码如 GBK、ISO-8859-1(西欧)、EUC-JP(日文)等
    - 若不正确解码，中文字符等会显示为乱码
2. Content-Type头部指定了字符集
    - 如 `Content-Type: text/html; charset=gbk`
```go
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
    
    // 当编码不是UTF-8时需要解码转换
    //引入包:golang.org/x/text/transform，解码  
    bodyReader := transform.NewReader(bufReader, res.NewDecoder())  
    content, _ := io.ReadAll(bodyReader)  
    fmt.Println(string(content))  
}
```
## Post
提交表单
```go
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
```
提交json格式内容
```go
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
```
提交文件
- <mark style="background: #FFF3A3A6;">muiltipart</mark>实现拼接请求
- multipart会随机创建一个boudnry
```go
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

```
---
## Download
实现带进度条的下载
```go
  
func downloadWithProcess(url, filename string) {  
    r, err := http.Get(url)  
    if err != nil {  
       fmt.Println(err)  
       return  
    }  
    defer func(Body io.ReadCloser) {  
       err := Body.Close()  
       if err != nil {  
          fmt.Println(err)  
       }  
    }(r.Body)  
    out, err := os.Create(filename)  
    if err != nil {  
       fmt.Println(err)  
       return  
    }  
    defer func(out *os.File) {  
       err := out.Close()  
       if err != nil {  
          fmt.Println(err)  
       }  
    }(out)  
    newReader := &Reader{  
       Reader: r.Body,  
       Total:  r.ContentLength,  
    }  
    _, err = io.Copy(out, newReader)  
    if err != nil {  
       fmt.Println(err)  
    }  
    fmt.Println("\nDone")  
  
}  
  
// Reader 重写Reader实现进度条功能  
// 本质是io.copy内部有循环调用Read，而我们重写了Read，才能实现进度条刷新  
type Reader struct {  
    io.Reader  
    Total   int64  
    Current int64  
}  
  
func (r *Reader) Read(p []byte) (n int, err error) {  
    n, err = r.Reader.Read(p)  
    r.Current += int64(n)  
    // \r实现每次打印都回到行首  
    fmt.Printf("\rDownloading: %.2f of %%100", float64(r.Current*10000/r.Total)/100)  
    return  
}
```

---
## Redirect
> DefaultClient是包内提供的一个已经设置了一些默认属性和Client变量
> 可以如下根据需求自定义client
```go
//默认的检查重定向实现
func defaultCheckRedirect(req *Request, via []*Request) error {  
  if len(via) >= 10 {  
   return errors.New("stopped after 10 redirects")  
  }  
  return nil  
}
```
设置重定向上限
```go
func redirectLimits() {  
    client := &http.Client{  
       CheckRedirect: func(r *http.Request, via []*http.Request) error {  
          if len(via) >= 10 {  
             //如果重定向次数过多  
             return errors.New("stopped after 10 redirects")  
          }  
          return nil  
       },  
    }  
    //重定向11次，报错  
    resp, err := client.Get("https://httpbin.org/absolute-redirect/11")  
    if err != nil {  
       fmt.Println(err)  
       return  
    }  
    defer resp.Body.Close()  
    content, _ := io.ReadAll(resp.Body)  
    fmt.Println(string(content))  
}
```
禁止重定向
```go

```
---
## Cookie

![[截屏2025-04-19 16.28.32.png|541x295]]

cookie 的分类有两种 一种是会话期 cookie 一种是持久性 cookie

```go
// 不用jar手动附加cookie  
func redirWithCookieManual() {  
    client := &http.Client{  
       CheckRedirect: func(r *http.Request, via []*http.Request) error {  
          //禁止重定向  
          return http.ErrUseLastResponse  
       },  
    }  
    firstReq, _ := http.NewRequest(http.MethodGet, "https://httpbin.org/cookies/set?freeform=sada&name=tie", nil)  
    firstRes, _ := client.Do(firstReq)  
    defer firstRes.Body.Close()  
    //获得服务端返回的cookies  
    //手动重定向地址  
    secondReq, _ := http.NewRequest(http.MethodGet, "https://httpbin.org/cookies", nil)  
    for _, cookie := range firstRes.Cookies() {  
       //将cookie附加到重定向的请求中  
       secondReq.AddCookie(cookie)  
    }  
    secondRes, _ := client.Do(secondReq)  
    defer secondRes.Body.Close()  
    content, _ := io.ReadAll(secondRes.Body)  
    fmt.Println(string(content))  
}
```
使用jar自动附加cookie
```go
// 使用Jar附加cookie  
func jarCookie() {  
    jar, _ := cookiejar.New(nil)  
    client := &http.Client{  
       Jar: jar,  
    }  
    r, _ := client.Get("https://httpbin.org/cookies/set?freeform=sada&name=tie")  
    defer r.Body.Close()  
    _, _ = io.Copy(os.Stdout, r.Body)  
}
```
>标准库只提供了会话期cookie
---
## Proxy
```go
proxyUrl, _ := url.Parse("http://127.0.0.1:7897")  //代理启动的端口
t := http.Transport{  
    Proxy: http.ProxyURL(proxyUrl),  
}  
//代理一般分两种，http代理和shadowsocks的代理,socks5  
client := &http.Client{Transport: &t}  
r, _ := client.Get("https://google.com")  
defer r.Body.Close()  
_, _ = io.Copy(os.Stdout, r.Body)
```
