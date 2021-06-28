package school

import (
	"github.com/gin-gonic/gin"
	"schoolServer/models"
	"schoolServer/pkg/app"
	"schoolServer/pkg/e"
)

// @Summary Get school
// @Produce  json
// @Param is_firstuniversity query string false "is_firstuniversity"
// @Param is_firstdiscipline query string false "is_firstdiscipline"
// @Param is_985 query string false "is_985"
// @Param is_211 query string false "is_211"
// @Param school_zone query string false "school_zone"
// @Param searchText query string false "searchText"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /school [get]
func GetSchool(c *gin.Context)  {
	//前端传值是"1"表示true、"0"表示false
	is_firstuniversity :=  false
	if arg := c.Query("is_firstuniversity"); arg != "" {
		if arg == "1" {
			is_firstuniversity = true
		}
	}
	is_firstdiscipline := false
	if arg := c.Query("is_firstdiscipline"); arg != "" {
		if arg == "1" {
			is_firstdiscipline = true
		}
	}
	is_985 := false
	if arg := c.Query("is_985"); arg != "" {
		if arg == "1" {
			is_985 = true
		}
	}
	is_211 := false
	if arg := c.Query("is_211"); arg != "" {
		if arg == "1" {
			is_211 = true
		}
	}
	school_zone := ""
	if arg := c.Query("school_zone"); arg != "" {
		school_zone = arg
	}
	searchText := ""
	if arg := c.Query("searchText"); arg != "" {
		searchText = arg
	}
	schoolParam := map[string]interface{}{
		"is_firstuniversity": is_firstuniversity,
		"is_firstdiscipline": is_firstdiscipline,
		"is_985": is_985,
		"is_211": is_211,
		"school_zone": school_zone,
		"searchText": searchText,
	}
	err, info := models.GetSchool(schoolParam)
	if err != nil {
		app.Error(c,e.ERROR,err,err.Error())
		return
	}
	app.OK(c, info, "OK")
}