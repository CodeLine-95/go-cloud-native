package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/CodeLine-95/go-cloud-native/pkg/logz"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils/id"
	"github.com/CodeLine-95/go-cloud-native/pkg/utils/traceId"
	"github.com/CodeLine-95/go-cloud-native/pkg/xlog"
	"github.com/gin-gonic/gin"
	"io"
	"strings"
	"time"
)

const maxMemory = 32 << 20

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId.SetTraceId(c, fmt.Sprintf("%v", id.Make.Make()))
		start := time.Now()
		params := c.Request.URL.RawQuery
		if c.Request.Method == "POST" {
			contentType := strings.Split(c.Request.Header.Get("Content-Type"), ";")[0]
			switch contentType {
			case "application/x-www-form-urlencoded":
				if err := c.Request.ParseForm(); err == nil {
					values := c.Request.PostForm
					jsonByte, _ := json.Marshal(values)
					params = string(jsonByte)
				}
			case "application/form-data":
				if err := c.Request.ParseMultipartForm(maxMemory); err == nil {
					values := c.Request.PostForm
					jsonByte, _ := json.Marshal(values)
					params = string(jsonByte)
				}
			case "multipart/form-data":
				if err := c.Request.ParseMultipartForm(maxMemory); err == nil {
					values := c.Request.PostForm
					jsonByte, _ := json.Marshal(values)
					params = string(jsonByte)
				}
			default:
				if requestBody, err := io.ReadAll(c.Request.Body); err == nil {
					c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
					params = string(requestBody)
				}
			}
		}

		// Stop timer
		end := time.Now()
		latency := end.Sub(start).Seconds() * 1e3

		xlog.Info(traceId.GetLogContext(c, c.Errors.ByType(gin.ErrorTypePrivate).String(),
			logz.F("params", params),
			logz.F("statusCode", c.Writer.Status()),
			logz.F("latency", latency),
			logz.F("bodySize", c.Writer.Size()),
		))
	}
}
