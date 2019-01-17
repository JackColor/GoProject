/*
*Auth: JackColor
*Date: 2019/1/13 下午4:35
*/
package logging

import (
	"fmt"
	"os"
	"strconv"
)

type LogFile struct {
	Path   string
	Name   string
	Level  int
	File   *os.File
	Error  *os.File
	LogMsg chan *LogMessage
}

func NewLogFile(config map[string]string) (Logger LoggerInterface, err error) {

	path, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("cant find the log path from the configer")
	}

	name, ok := config["log_name"]

	if !ok {
		err = fmt.Errorf("cant find the log name from the configer")
	}

	level, ok := config["log_level"]
	if !ok {

		level = "DEBUG"
	}

	levelNum := getLogLevelNum(level)

	chanSizeStr, ok := config["chan_size"]
	if !ok {
		chanSizeStr = "50000"
	}

	chanSize, err := strconv.Atoi(chanSizeStr)
	if err != nil {
		chanSize = 50000
	}

	Logger = &LogFile{
		Path:   path,
		Name:   name,
		Level:  levelNum,
		LogMsg: make(chan *LogMessage, chanSize),
	}

	Logger.Init()
	//fmt.Println(Logger)
	return

}

func (f *LogFile) Init() {

	fileName := fmt.Sprintf("%s/%s.log", f.Path, f.Name)
	NormalFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic(fmt.Sprintf("open normal file name (%s) failed errors:%s", fileName, err))
	}

	f.File = NormalFile

	errorName := fmt.Sprintf("%s/%s.error.log", f.Path, f.Name)
	ErrorFile, err := os.OpenFile(errorName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic(fmt.Sprintf("open normal file name (%s) failed errors:%s", errorName, err))

	}
	f.Error = ErrorFile
	go f.writeLogInfo()
	return
}

func (f *LogFile) SetLevel(level int) {
	if level < LogDebug || level > LogFatal {
		level = LogInfo
	}

	f.Level = level

}

func (f *LogFile) writeLogInfo() {
	for message := range f.LogMsg {
		file := f.File
		if message.ErrorFatal {
			file = f.Error
		}

		fmt.Fprintf(file, message.Msg)

	}

}

func (f *LogFile) LogDebug(format string, args ... interface{}) {

	if f.Level > LogDebug {
		return
	}

	// 加上 日志 级别 时间 函数 函数行号 msg 信息
	// 2019-01-01 01:01:01:9999 DEBUG -FILENAME-FUNC:LINE : MSG
	//fmt.Fprintf(f.File, format, args...)

	LogMsg, err := getLogMsgFomat(LogDebug, format, args...)
	if err != nil {
		return
	}

	select {
	case f.LogMsg <- LogMsg:

	default:


	}

}

func (f *LogFile) LogTrace(format string, args ... interface{}) {

	if f.Level > LogTrace {
		return
	}

	LogMsg, err := getLogMsgFomat(LogTrace, format, args...)
	if err != nil {
		return
	}

	select {
	case f.LogMsg <- LogMsg:

	default:


	}

}

func (f *LogFile) LogInfo(format string, args ... interface{}) {
	if f.Level > LogInfo {
		return
	}

	LogMsg, err := getLogMsgFomat(LogInfo, format, args...)
	if err != nil {
		return
	}

	select {
	case f.LogMsg <- LogMsg:

	default:


	}
}

func (f *LogFile) LogWarn(format string, args ... interface{}) {
	if f.Level > LogWarn {
		return
	}

	LogMsg, err := getLogMsgFomat(LogWarn, format, args...)
	if err != nil {
		return
	}

	select {
	case f.LogMsg <- LogMsg:

	default:


	}
}
func (f *LogFile) LogErrors(format string, args ... interface{}) {

	//fmt.Println("errors msg start.....")

	if f.Level > LogError {
		return
	}
	LogMsg, err := getLogMsgFomat(LogError, format, args...)
	if err != nil {
		return
	}
	//fmt.Println(LogMsg)
	select {
	case f.LogMsg <- LogMsg:

	default:


	}
}

func (f *LogFile) LogFatal(format string, args ... interface{}) {
	if f.Level > LogFatal {
		return
	}
	LogMsg, err := getLogMsgFomat(LogFatal, format, args...)
	if err != nil {
		return
	}

	select {
	case f.LogMsg <- LogMsg:

	default:


	}
}
