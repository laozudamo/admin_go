package controller

import (
	"admin_go/response"
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	// 单文件
	file, err := c.FormFile("file")
	if err != nil {
		response.Err(c, 200, 400, "文件错误", nil)
	}
	fileName := fmt.Sprintf("%d%s", time.Now().Unix(), filepath.Ext(file.Filename))
	// 上传文件至指定的完整文件路径
	c.SaveUploadedFile(file, "./uploads/"+fileName)
	imageUrl := "http://localhost:8080/uploads/" + fileName
	response.Success(c, 200, "上传成功", imageUrl)
}
