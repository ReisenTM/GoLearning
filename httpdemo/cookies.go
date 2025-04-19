package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
)

// 不用jar自动附加cookie
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

//func main() {
//	jarCookie()
//}
