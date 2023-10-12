package response

import (
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/CodeLine-95/go-cloud-native/tools/traceId"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Default = &response{}

// Error 失败数据处理
func Error(c *gin.Context, code int, err error, msg string) {
	res := Default.Clone()
	res.SetInfo(msg)
	if err != nil {
		res.SetInfo(err.Error())
	}
	if msg != "" {
		res.SetMsg(msg)
	}
	res.SetTraceID(traceId.GetTraceId(c))
	res.SetCode(int32(code))
	res.SetSuccess(false)
	// 记录日志
	xlog.Error(traceId.GetLogContext(c, msg, logz.F("err", err), logz.F("response", res)))
	// 写入上下文
	c.Set("result", res)
	// 返回结果集
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// OK 通常成功数据处理
func OK(c *gin.Context, data any, msg string) {
	res := Default.Clone()
	res.SetData(data)
	res.SetSuccess(true)
	if msg != "" {
		res.SetMsg(msg)
		res.SetInfo(msg)
	}
	res.SetTraceID(traceId.GetTraceId(c))
	res.SetCode(http.StatusOK)
	// 记录日志
	xlog.Info(traceId.GetLogContext(c, msg, logz.F("response", res)))
	// 写入上下文
	c.Set("result", res)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// PageOK 分页数据处理
func PageOK(c *gin.Context, result any, count int, pageIndex int, pageSize int, msg string) {
	var res page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}

// Custum 兼容函数
func Custum(c *gin.Context, data gin.H) {
	data["requestId"] = traceId.GetTraceId(c)
	c.Set("result", data)
	c.AbortWithStatusJSON(http.StatusOK, data)
}
