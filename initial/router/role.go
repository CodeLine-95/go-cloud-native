package router

import (
	"github.com/CodeLine-95/go-cloud-native/internal/app/logic/role"
	"github.com/gin-gonic/gin"
)

func Role(r *gin.RouterGroup) {
	c := r.Group("/role")
	c.POST("/list", role.List)
	c.POST("/add", role.Add)
	c.POST("/edit", role.Edit)
	c.POST("/del", role.Del)
}
