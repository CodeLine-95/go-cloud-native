package xlog

import (
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils/id"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils/traceId"
	"github.com/CodeLine-95/go-cloud-native/pkg/xlog"
	logz "github.com/CodeLine-95/go-tools"
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId.SetTraceId(c, fmt.Sprintf("%v", id.Make.Make()))
		c.Next()
		xlog.Debug(traceId.GetLogContext(c, c.Errors.String(), logz.F("reqData", c.Params)))
	}
}
