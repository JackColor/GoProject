/*
*Auth: JackColor
*Date: 2019/1/12 下午10:57
*/
package handler

import "fmt"

var logger LogInterface

func InitLogger(name string, config map[string]string) (err error) {

	switch name {
	case "file":
		logger, err = NewFileLog(config)
		if err != nil {
			return
		}
	case "console":
		logger, err = NewConsoleLog(config)
		if err != nil {
			return
		}
		// console
	default:
		err = fmt.Errorf("cant support the log type")
		return
	}

	return
}

func DebugLog(format string, args ...interface{}) {

	logger.Debug(format, args...)

}

func InfoLog(format string, args ...interface{}) {

	logger.Info(format, args...)

}

func WarnLog(format string, args ...interface{}) {

	logger.Warn(format, args...)

}

func TraceLog(format string, args ...interface{}) {

	logger.Trace(format, args...)

}

func ErrorsLog(format string, args ...interface{}) {

	logger.Error(format, args...)

}

func FatalLog(format string, args ...interface{}) {

	logger.Fatal(format, args...)

}
