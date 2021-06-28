package logger

import (
	"schoolServer/pkg/setting"
	"fmt"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return fmt.Sprintf("%s%s", setting.LogSetting.RuntimeRootPath, setting.LogSetting.LogSavePath)
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		setting.LogSetting.LogSaveName,
		time.Now().Format(setting.LogSetting.TimeFormat),
		setting.LogSetting.LogFileExt,
	)
}
