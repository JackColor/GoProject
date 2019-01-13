/*
*Auth: JackColor
*Date: 2019/1/12 下午7:49
*/
package handler

type LogInterface interface {
	Init()
	SetLevel(level int)
	Debug(format string, args ...interface{})
	Trace(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}


