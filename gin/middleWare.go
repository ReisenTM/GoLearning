package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in")
	//go funcXXX(c.Copy()) 注意在中间件使用go routine一定要传c的拷贝，不然线程不安全
	c.Next() //执行下一个中间键
	//c.Abort()禁止后续所有中间件
	fmt.Println("m1 out")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in")

	// 跨中间件传值
	c.Set("name", "帕鲁")
	c.Next()
	fmt.Println("m2 out")
}

func getHandler(c *gin.Context) {
	name, ok := c.Get("name")
	if !ok {
		fmt.Println("匿名用户")
		name = ""
	}
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func mwTest() {
	r := gin.Default()
	//r.Use(m1,m2,getHandler)//全局中间件设置，生效于所有请求
	//用处：登录鉴权等等
	r.GET("/", m1, m2, getHandler)
	r.Run(":8080")
}
