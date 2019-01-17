/*
*Auth: JackColor
*Date: 2019/1/13 下午4:35
*/
package logging

type LoggerInterface interface {
	Init()
	SetLevel(level int)
	LogDebug(format string,args... interface{})
	LogTrace(format string,args... interface{})
	LogInfo(format string,args... interface{})
	LogWarn(format string,args... interface{})
	//errors 文件
	LogErrors(format string,args... interface{})
	LogFatal(format string,args... interface{})
}

