package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func download(url, filename string) {
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
	_, err = io.Copy(out, r.Body)
	if err != nil {
		fmt.Println(err)
	}
}
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

//func main() {
//	url := "https://i1.hdslb.com/bfs/face/0d2f52b9e1dae83b2cd710b38ee911533db32774.jpg"
//	filename := "test.jpg"
//	downloadWithProcess(url, filename)
//}
