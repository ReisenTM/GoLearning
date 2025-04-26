package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func form() {
	r := gin.Default()
	r.LoadHTMLFiles("form/login.html", "form/done.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/done", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.HTML(http.StatusOK, "done.html", gin.H{
			"Name": username,
			"Pass": password,
		})
	})
	r.Run(":8080")
}
