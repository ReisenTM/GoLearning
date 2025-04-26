package main

import "github.com/gin-gonic/gin"

// 获取QueryString
// .../query=张三
func QueryString() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		//通过请求获取query string参数,多个请求用&连接
		name := c.Query("name")
		age := c.Query("age")
		//name:=c.DefaultQuery("name", "Tom")  取不到就用默认值
		//name,ok :=c.GetQuery("name") 返回两个参数 string,bool
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":8080")
}
