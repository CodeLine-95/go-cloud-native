package routers

import "github.com/gin-gonic/gin"

func Login(r *gin.RouterGroup) {
	loginG := r.Group("/base")
	loginG.POST("/login")
}
