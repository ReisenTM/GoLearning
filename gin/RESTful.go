package main

import "github.com/gin-gonic/gin"

func RESTfulExample() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})
	r.POST("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})
	r.PUT("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})
	r.DELETE("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	err := r.Run()
	if err != nil {
		return
	}
}
