package log

import (
	"go.uber.org/zap"
	global "go_xx/config"
)

func InitLogin(fileName string, maxSize, maxBackup, maxAge int, compress bool, logType, level string) {

}

// LogIf err级别
func LogIf(err error) {
	if err != nil {
		global.Logger.Error("Error occurred:", zap.Error(err))
	}
}

// LogInfoIf info级别
func LogInfoIf(err error) {
	if err != nil {
		global.Logger.Info("Error occurred:", zap.Error(err))
	}
}

// LogWarnIf warn级别
func LogWarnIf(err error) {
	if err != nil {
		global.Logger.Warn("Error occurred:", zap.Error(err))
	}
}

func Info(moduleName string, fields ...zap.Field) {
	global.Logger.Info(moduleName, fields...)
}
