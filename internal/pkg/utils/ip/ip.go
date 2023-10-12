package ip

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// getIP 获取用户真实ip
func ClientIP(c *gin.Context) string {
	ip := c.Request.RemoteAddr
	realIP := strings.TrimSpace(c.GetHeader("X-Real-Ip"))
	if realIP != "" {
		ip = realIP
	}
	return ip
}
