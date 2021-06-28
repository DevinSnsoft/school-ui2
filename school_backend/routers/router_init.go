package routers

import (
	"github.com/gin-gonic/gin"
	middlewares "schoolServer/middleWare"
	_ "schoolServer/models"
)

//无需认证
func sysNoCheckRoleRouter(r *gin.RouterGroup) {
	SysBaseRouter(r)
	SchoolRouter(r)
}

func sysCheckRoleRouter(r *gin.RouterGroup, authMiddle gin.HandlerFunc) {
	SysBaseRouterWithMiddleWare(r, authMiddle)
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middlewares.Cors())

	g := r.Group("")
	//jwt := middleWare.JWT()

	sysNoCheckRoleRouter(g)
	//sysCheckRoleRouter(g,jwt)
	return r
}
