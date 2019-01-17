/*
*Auth: JackColor
*Date: 2019/1/13 下午5:58
*/
package logging

var LogInstance LoggerInterface

func Init(name string, conf map[string]string) (err error) {

	switch name {
	case "file":
		LogInstance, err = NewLogFile(conf)
		if err != nil {
			return err
		}
	case "console":
		LogInstance, err = NewLogConsole(conf)

	}

	return

}

func Debug(format string, args ...interface{}) {

	LogInstance.LogDebug(format, args...)

}

func Trace(format string, args ...interface{}) {

	LogInstance.LogTrace(format, args...)

}

func Info(format string, args ...interface{}) {

	LogInstance.LogInfo(format, args...)

}

func Warn(format string, args ...interface{}) {

	LogInstance.LogWarn(format, args...)

}

func Error(format string, args ...interface{}) {

	LogInstance.LogErrors(format, args...)

}

func Fatal(format string, args ...interface{}) {

	LogInstance.LogFatal(format, args...)

}
