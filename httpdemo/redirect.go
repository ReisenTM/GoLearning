package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

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
func redirectForbidden() {
	client := &http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Get("https://httpbin.org/cookies/set?freeform=sadasd")
	resp2, err := http.DefaultClient.Get("https://httpbin.org/cookies/set?freeform=sadasd")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	defer resp2.Body.Close()
	fmt.Println("禁止重定向地址", resp.Request.URL.String())
	fmt.Println("原本重定向地址", resp2.Request.URL.String())
}

//func main() {
//	//重定向
//	//返回状态码 ,3xx
//	redirectForbidden()
//}
