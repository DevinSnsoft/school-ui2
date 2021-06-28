package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "schoolServer/docs"
	"schoolServer/models"
	"schoolServer/pkg/gredis"
	"schoolServer/pkg/logger"
	"schoolServer/pkg/setting"
	"schoolServer/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	gredis.Setup()
	logger.Setup()
}
// @title schoolServer
// @version 1.0
// @description base on gin
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	routersInit := routers.InitRouter()
	routersInit.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routersInit.Run(endPoint)
}