package common

import (
	"github.com/CodeLine-95/go-cloud-native/common/constant"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils/resp"
	"github.com/gin-gonic/gin"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorParams,
			ErrMsg:  constant.ErrorMsg[constant.ErrorParams],
		}, nil)
	}
	// 获取文件后缀名
	fileExt := path.Ext(file.Filename)

	// 新文件名
	fileName := utils.MD5(file.Filename+strconv.FormatInt(time.Now().Unix(), 10)) + fileExt

	// 目录转换
	filePath := filepath.Join(utils.NewMkdir("uploads"), "/", fileName)

	// 保存文件
	if saveErr := c.SaveUploadedFile(file, filePath); saveErr != nil {
		resp.Response(c, &constant.Error{
			ErrCode: constant.ErrorUploadImage,
			ErrMsg:  "文件上传失败",
		}, nil)
	}

	resp.Response(c, &constant.Error{
		ErrCode: constant.Success,
		ErrMsg:  constant.ErrorMsg[constant.Success],
	}, gin.H{
		"fileUrl": filePath,
	})
}
