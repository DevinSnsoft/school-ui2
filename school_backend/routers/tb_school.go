package routers

import (
	"github.com/gin-gonic/gin"
	"schoolServer/apis/school"
)

func SchoolRouter(r *gin.RouterGroup)  {
	r.GET("/school", school.GetSchool)
}
