package router

import (
	"github.com/CodeLine-95/go-cloud-native/internal/app/logic"
	"github.com/gin-gonic/gin"
)

func Common(r *gin.RouterGroup) {
	c := r.Group("/base")
	c.POST("/upload", logic.Upload)

	a := r.Group("/auth")
	a.POST("/token", logic.Login)
}
