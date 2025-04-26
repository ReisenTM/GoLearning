package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func singleFile() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(file)
		dst := path.Join("./", file.Filename)
		err = c.SaveUploadedFile(file, dst)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": file})
	})
	r.Run(":8080")
}
func multiFile() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		for _, files := range file.File {
			for _, file := range files {
				fileName := file.Filename
				dst := path.Join("./uploadFiles", fileName)
				err = c.SaveUploadedFile(file, dst)
			}
		}
		c.JSON(http.StatusOK, gin.H{"upload": "complete"})
	})
	r.Run(":8080")
}
