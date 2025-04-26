package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func router() {
	r := gin.Default()

	//r.Any("/", func(c *gin.Context) {}) //匹配任何请求

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"404":     "not found",
			"message": "啥都木有",
		})
	})
	//路由组，提取前缀相同的请求
	//路由组支持嵌套
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "index",
			})
		})
		videoGroup.POST("/upload", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "upload",
			})
		})
		videoGroup.POST("/delete", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "delete",
			})
		})
	}
	r.Run(":8080")
}
