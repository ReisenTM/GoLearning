package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func redir() {
	r := gin.Default()
	//请求重定向 ->两次请求
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com") //返回重定向状态，设置重定向路径
	})
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c) //重新载入
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		return
	}
	//请求转发 ->一次请求

}
