package routers

import (
	"github.com/CodeLine-95/go-cloud-native/controllers/common"
	"github.com/gin-gonic/gin"
)

func Common(r *gin.RouterGroup) {
	c := r.Group("/common")
	c.POST("/upload", common.Upload)

	a := r.Group("/auth")
	a.POST("/login", common.Login)
}
