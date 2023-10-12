package logic

import (
	"github.com/CodeLine-95/go-cloud-native/internal/app/constant"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/response"
	"github.com/CodeLine-95/go-cloud-native/tools/utils"
	"github.com/gin-gonic/gin"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		response.Error(c, constant.ErrorParams, err, constant.ErrorMsg[constant.ErrorParams])
		return
	}
	// 获取文件后缀名
	fileExt := path.Ext(file.Filename)

	// 新文件名
	fileName := utils.MD5(file.Filename+strconv.FormatInt(time.Now().Unix(), 10)) + fileExt

	// 目录转换
	filePath := filepath.Join(utils.NewMkdir("assets/uploads"), "/", fileName)

	// 保存文件
	if saveErr := c.SaveUploadedFile(file, filePath); saveErr != nil {
		response.Error(c, constant.ErrorUploadImage, saveErr, "文件上传失败")
		return
	}

	response.OK(c, gin.H{
		"fileUrl": filePath,
	}, constant.ErrorMsg[constant.Success])
}
