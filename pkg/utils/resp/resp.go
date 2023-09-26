package resp

import (
	"github.com/CodeLine-95/go-cloud-native/common/constant"
	"github.com/CodeLine-95/go-cloud-native/pkg/logz"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils/traceId"
	"github.com/CodeLine-95/go-cloud-native/pkg/xlog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, codeErr *constant.Error, data any) {
	traceID := traceId.GetTraceId(c)
	err := constant.ErrorSuccess
	if codeErr != nil {
		err = codeErr
	}
	msg := err.ErrMsg
	xlog.Error(traceId.GetLogContext(
		c,
		msg,
		logz.F("err", err),
	))
	c.JSON(http.StatusOK, constant.Response{
		ErrNo:   err.ErrCode,
		Msg:     msg,
		Info:    "",
		TraceID: traceID,
		Data:    data,
	})
}
