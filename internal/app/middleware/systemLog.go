package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	common "github.com/CodeLine-95/go-cloud-native/common/models"
	"github.com/CodeLine-95/go-cloud-native/internal/app/models"
	"github.com/CodeLine-95/go-cloud-native/internal/app/service/logs"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/jwtToken"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/utils/ip"
	"github.com/CodeLine-95/go-cloud-native/internal/pkg/xlog"
	"github.com/CodeLine-95/go-cloud-native/tools/id"
	"github.com/CodeLine-95/go-cloud-native/tools/logz"
	"github.com/CodeLine-95/go-cloud-native/tools/traceId"
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
			case "application/json":
				if requestBody, err := io.ReadAll(c.Request.Body); err == nil {
					c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
					jsonStruct := make(map[string]string, 0)
					_ = json.Unmarshal(requestBody, &jsonStruct)
					jsonByte, _ := json.Marshal(jsonStruct)
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

		token := jwtToken.GetToken(c.Request, "")

		cloudLog := &models.CloudLog{
			LogID:      traceId.GetTraceId(c),
			LogName:    "--",
			RequestUrl: c.Request.RequestURI,
			Method:     c.Request.Method,
			ClientIP:   ip.ClientIP(c),
			Level:      "info",
			AppType:    0,
			ParamsData: params,
			ModelTime: common.ModelTime{
				CreateTime: uint32(time.Now().Unix()),
			},
		}

		_ = logs.SaveData(cloudLog, token)
	}
}
