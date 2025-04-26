package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//ShouldBInd可以根据请求的类型自动绑定

type UserInfo struct {
	Username string `form:"username"` //form用来和请求里的key对应
	Password string `form:"password"`
}

func bind() {
	r := gin.Default()
	r.POST("/query", func(c *gin.Context) {
		var user UserInfo
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"data": user})
		}
	})

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
