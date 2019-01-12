/*
*Auth: JackColor
*Date: 2019/1/12 下午7:47
*/
package handler

//静态字段
const (
	Debug = iota
	Trace
	Info
	Warn
	Error
	Fatal
)

func getLevelText(level int) string {

	switch level {
	case Debug:
		return "Debug"
	case Warn:
		return "Warn"
	case Trace:
		return "Trace"
	case Fatal:
		return "Fatal"
	case Info:
		return "Info"
	case Error:
		return "Error"

	}

	return "UNKNOWN"

}
