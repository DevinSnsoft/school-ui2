package upload

import (
	"fmt"
	"os"
	"path"
	"schoolServer/pkg/file"
	"schoolServer/pkg/setting"
	"schoolServer/pkg/utils"
	"strings"
)

// GetVideoFullUrl get the full access path
func GetVideoFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetVideoPath() + name
}

// GetVideoName get image name
func GetVideoName(name string) string {
	ext := path.Ext(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = utils.EncodeMD5(fileName)

	return fileName + ext
}

// GetImagePath get save path
func GetVideoPath() string {
	return setting.AppSetting.VideoSavePath
}

// GetImageFullPath get full save path
func GetVideoFullPath() string {
	return setting.LogSetting.RuntimeRootPath + GetVideoPath()
}

// CheckVideo check if the file exists
func CheckVideo(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}