package routers

import (
	"github.com/gin-gonic/gin"
	"schoolServer/handler"
)

func SysBaseRouter(r *gin.RouterGroup) {
	r.GET("/ping", handler.Ping)
}

func SysBaseRouterWithMiddleWare(r *gin.RouterGroup, middle gin.HandlerFunc )  {
	r.Use(middle).GET("/ping", handler.Ping)
}