package models

type school struct{
	School_id int `gorm:"column:school_id"`
	School_serial string `gorm:"column:school_serial"`
	School_code string `gorm:"column:school_code"`
	School_name string `gorm:"column:school_name"`
	School_zone string `gorm:"column:school_zone"`
	School_city string `gorm:"column:school_city"`
	School_department string `gorm:"column:school_department"`
	School_level string `gorm:"column:school_level"`
	School_type string `gorm:"column:school_type"`
	Is_firstuniversity bool `gorm:"column:is_firstuniversity"`
	Is_firstdiscipline bool `gorm:"column:is_firstdiscipline"`
	Is_985 bool `gorm:"column:is_985"`
	Is_211 bool `gorm:"column:is_211"`
	Is_del bool `gorm:"column:is_del"`
}

func GetSchool(schoolParam map[string]interface{}) (error, []school) {
	var schooldata []school
	is_firstuniversity := schoolParam["is_firstuniversity"].(bool)
	is_firstdiscipline := schoolParam["is_firstdiscipline"].(bool)
	is_985 := schoolParam["is_985"].(bool)
	is_211 := schoolParam["is_211"].(bool)
	school_zone := schoolParam["school_zone"].(string)
	searchText := schoolParam["searchText"].(string)

	//err := db.Table("school").Where("is_firstuniversity = ? and is_firstdiscipline = ? and is_985 = ? and is_211 = ? and (school_zone = ? or ? = '') and school_name like ? and is_del = false",
	//	is_firstuniversity, is_firstdiscipline, is_985, is_211, school_zone, school_zone, "%"+searchText+"%").Find(&schooldata).Error

	db := db.Where("is_del = false")
	if is_firstuniversity {
		db = db.Where("is_firstuniversity = ?", is_firstuniversity)
	}
	if is_firstdiscipline {
		db = db.Where("is_firstdiscipline = ?", is_firstdiscipline)
	}
	if is_985 {
		db = db.Where("is_985 = ?", is_985)
	}
	if is_211 {
		db = db.Where("is_211 = ?", is_211)
	}
	if len(school_zone) > 0 {
		db = db.Where("school_zone = ?", school_zone)
	}
	if len(searchText) > 0 {
		db = db.Where("school_name like ?", "%"+searchText+"%")
	}
	err := db.Table("school").Find(&schooldata).Error



	//var err error
	//if is_firstuniversity {
	//	err = db.Table("school").Where("is_firstuniversity = ?").Find(&schooldata).Error
	//}
	//if is_firstdiscipline {
	//	err = db.Table("school").Where("is_firstdiscipline = ?").Find(&schooldata).Error
	//}
	//if is_985 {
	//	err = db.Table("school").Where("is_985 = ?").Find(&schooldata).Error
	//}
	//if is_211 {
	//	err = db.Table("school").Where("is_211 = ?").Find(&schooldata).Error
	//}
	//if school_zone {
	//	err = db.Table("school").Where("school_zone").Find(&schooldata).Error
	//}
	//if searchText {
	//	err = db.Table("school").Where("%"+searchText+"%").Find(&schooldata).Error
	//}


	return err, schooldata
}
//若本地安装有redis用以下版本，解开这个GetSchool注释上一个GetSchool
//func GetSchool(schoolParam map[string]interface{}) (error, []school) {
//	var schooldata []school
//	is_firstuniversity := schoolParam["is_firstuniversity"].(bool)
//	is_firstdiscipline := schoolParam["is_firstdiscipline"].(bool)
//	is_985 := schoolParam["is_985"].(bool)
//	is_211 := schoolParam["is_211"].(bool)
//	school_zone := schoolParam["school_zone"].(string)
//	searchText := schoolParam["searchText"].(string)
//
//	exist, err := gredis.RedisConn.Exists(e.AllSchoolRedisKey)
//	if err != nil {
//		logger.Error(err.Error())
//	}
//	if exist {
//		gredis.RedisConn.GetObject(e.AllSchoolRedisKey, &schooldata)
//	} else {
//		err = db.Table("school").Where("is_del = false").Find(&schooldata).Error
//		if len(schooldata) != 0 {
//			gredis.RedisConn.Set(e.AllSchoolRedisKey, schooldata, 0)
//		}
//	}
//	var schooldatatemp []school
//	for k, _ := range schooldata {
//		if schooldata[k].Is_firstuniversity == is_firstuniversity && schooldata[k].Is_firstdiscipline == is_firstdiscipline && schooldata[k].Is_985 == is_985 && schooldata[k].Is_211 == is_211 && (schooldata[k].School_zone == school_zone || school_zone == "") && (strings.Contains(schooldata[k].School_name, searchText) || searchText == "") {
//			schooldatatemp = append(schooldatatemp, schooldata[k])
//		}
//	}
//	return err, schooldatatemp
//}