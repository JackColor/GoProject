/*
*Auth: JackColor
*Date: 2019/1/13 下午4:51
*/
package logging

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type LogMessage struct {
	Msg        string
	ErrorFatal bool
}

func getLogLevelNum(level string) int {

	switch level {
	case "DEBUG":
		return LogDebug
	case "TRACE":
		return LogTrace
	case "INFO":
		return LogInfo
	case "WARN":
		return LogWarn
	case "ERROR":
		return LogError
	case "FATAL":
		return LogFatal
	default:
		return LogInfo
	}

}

func getLogLevelText(level int) string {

	switch level {
	case LogDebug:
		return "DEBUG"
	case LogTrace:
		return "TRACE"
	case LogInfo:
		return "INFO"
	case LogWarn:
		return "WARN"
	case LogError:
		return "ERROR"
	case LogFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}

}

func getFuncInfo() (FileName, FuncName string, LineNo int, err error) {

	pc, FileName, LineNo, ok := runtime.Caller(3)
	if !ok {
		err = fmt.Errorf("get func info failed")
		return
	}

	FuncName = runtime.FuncForPC(pc).Name()
	FileName = path.Base(FileName)
	FuncName = path.Base(FuncName)
	return
}

func getLogMsgFomat(level int, format string, args ... interface{}) (LogMsg *LogMessage, err error) {
	// 2019-01-01 01:01:01:9999 DEBUG -FILENAME-FUNC:LINE : MSG
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05.9999")

	levelText := getLogLevelText(level)

	FileName, FuncName, LineNo, err := getFuncInfo()
	if err != nil {
		return
	}

	msg := fmt.Sprintf("%s %s -%s-%s:%d %s \n", timeStr, levelText, FileName, FuncName, LineNo, format)
	Msg := fmt.Sprintf(msg, args...)

	LogMsg = &LogMessage{Msg: Msg, ErrorFatal: false}

	if level == LogError || level == LogFatal {
		LogMsg.ErrorFatal = true
	}

	return

}
