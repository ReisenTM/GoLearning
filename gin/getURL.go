package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取请求的path（URI）参数
func getUrl() {
	r := gin.Default()
	r.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name") //返回值为string
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	r.Run(":8080")
}
