/*
*Auth: JackColor
*Date: 2019/1/12 下午8:58
*/
package handler

import (
	"fmt"
	"path"
	"runtime"

	"time"
)

type LogData struct {
	Msg string
	//DateStr  string
	//FileName string
	//FuncName string
	LogLevel string
	//LineNO   int
	ErrorAndFatal bool
}

func getLogLine() (FileName, FuncName string, Line int) {

	pc, filename, line, ok := runtime.Caller(4)

	if ok {
		FileName = filename
		FuncName = runtime.FuncForPC(pc).Name()
		Line = line
		return
	}
	fmt.Println("get log line info errors")
	return
}

func writeLog(level int, format string, args ...interface{}) *LogData {

	now := time.Now()
	// 获取时间 字符串
	nowStr := now.Format("2006-01-02 15:04:05.999")
	filename, funcname, line := getLogLine()

	filename = path.Base(filename)
	funcname = path.Base(funcname)

	levelText := getLevelText(level)

	// 时间 级别 文件名称 函数名 函数行号  格式字符串
	msg := fmt.Sprintf("%s %s [%s %s:%d] %s \n", nowStr, levelText, filename, funcname, line, format,)

	msg = fmt.Sprintf(msg,args...)


	MsgData := &LogData{
		Msg:           msg,
		ErrorAndFatal: false,
	}
	if level == Error || level == Fatal {

		MsgData.ErrorAndFatal = true

	}

	return MsgData
	//fmt.Fprintf(file, msg, args...)

}
