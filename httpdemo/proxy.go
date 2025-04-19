package main

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	proxyUrl, _ := url.Parse("http://127.0.0.1:7897")
	t := http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	//代理一般分两种，http代理和shadowsocks的代理,socks5
	client := &http.Client{Transport: &t}
	r, _ := client.Get("https://google.com")
	defer r.Body.Close()
	_, _ = io.Copy(os.Stdout, r.Body)
}
