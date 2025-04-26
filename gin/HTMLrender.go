package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin框架渲染html
func render() {
	r := gin.Default()
	//通过路径加载模板
	r.LoadHTMLGlob("templates/**/*")
	//r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	//当访问该url时，title处的内容会被渲染为指定内容
	r.GET("posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})
	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})
	r.Run(":8080")
}
