/*
*Auth: JackColor
*Date: 2019/1/13 下午4:35
*/
package logging

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type LogFile struct {
	Path        string
	Name        string
	Level       int
	LogType     string
	LogSplitFre int64
	Time        int
	File        *os.File
	Error       *os.File

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

	LogType, ok := config["log_type"]

	if !ok {
		LogType = "hour"
	}

	var LogSplitSize int

	if LogType == "size" {
		splitFre, ok := config["log_size"]
		if !ok {
			splitFre = "104857600"
		}

		LogSplitSize, err = strconv.Atoi(splitFre)
		if err != nil {
			LogSplitSize = 104857600
		}

	}

	Logger = &LogFile{
		Path:        path,
		Name:        name,
		Level:       levelNum,
		LogType:     LogType,
		LogSplitFre: int64(LogSplitSize),
		Time:        time.Now().Hour(),
		LogMsg:      make(chan *LogMessage, chanSize),
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

func (f *LogFile) TimeType() (err error) {

	Now := time.Now()
	currentHour := Now.Hour()
	if currentHour == f.Time {
		fmt.Println("time is same hahahahha")
		return
	}

	f.Time = currentHour

	//当前文件
	currentNormalFileName := fmt.Sprintf("%s/%s.log", f.Path, f.Name)
	// 备份文件

	NormalFileName := fmt.Sprintf("%s/%s.log_%d_%d", f.Path, f.Name, Now.Hour(), Now.Minute())

	currentErrorFileName := fmt.Sprintf("%s/%s.error.log", f.Path, f.Name)
	ErrorFileName := fmt.Sprintf("%s/%s.error.log.%d_%d", f.Path, f.Name, Now.Hour(), Now.Minute())

	err = f.File.Close() //关闭正常文件

	if err != nil {
		err = fmt.Errorf("close the file failed in Time Type when close the normal file,err:%s", err)

		return
	}
	err = f.Error.Close() //关闭 错误文件

	if err != nil {

		err = fmt.Errorf("close the file failed in Time Type when close the error file,err:%s", err)

		return
	}

	os.Rename(currentNormalFileName, NormalFileName)

	os.Rename(currentErrorFileName, ErrorFileName)

	//打开正常 文件
	NorFile, err := os.OpenFile(currentNormalFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		err = fmt.Errorf("open the file failed in Time Type when open the normal file,err:%s", err)
		return
	}

	// 打开错误文件
	ErrorFile, err := os.OpenFile(currentErrorFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		err = fmt.Errorf("open the file failed in Time Type when open the error file,err:%s", err)

		return
	}

	f.File = NorFile

	f.Error = ErrorFile

	return
}

func (f *LogFile) SizeType() (err error) {

	//文件大小分割
	//拿到 当前文件的大小
	FileInfo, err := f.File.Stat()

	if err != nil {
		err = fmt.Errorf("get the file size from the file failed err:%s", err)
		return
	}

	Now := time.Now()

	currentNormalFileSize := FileInfo.Size()

	if currentNormalFileSize < f.LogSplitFre {
		return
	} else {
		currentNormalFileName := fmt.Sprintf("%s/%s.log", f.Path, f.Name)
		NormalFileName := fmt.Sprintf("%s/%s.log_%d_%d_%d", f.Path, f.Name, Now.Hour(), Now.Minute(), Now.Second())

		f.File.Close()
		os.Rename(currentNormalFileName, NormalFileName)

		NorFile, Normalerr := os.OpenFile(currentNormalFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if Normalerr != nil {
			Normalerr = fmt.Errorf("open the file failed in Time Type when open the normal file,err:%s", err)
			err = Normalerr
			return
		}

		f.File = NorFile

	}

	ErrorFileInfo, err := f.Error.Stat()
	if err != nil {

		return
	}

	currentErrorFileSize := ErrorFileInfo.Size()

	if currentErrorFileSize < f.LogSplitFre {
		return
	} else {

		currentErrorFileName := fmt.Sprintf("%s/%s.error.log", f.Path, f.Name)
		ErrorFileName := fmt.Sprintf("%s/%s.error.log.%d_%d_%d", f.Path, f.Name, Now.Hour(), Now.Minute(), Now.Second())

		f.Error.Close()

		os.Rename(currentErrorFileName, ErrorFileName)

		ErrorFile, Errorerr := os.OpenFile(currentErrorFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)

		if Errorerr != nil {
			Errorerr = fmt.Errorf("open the file failed in Time Type when open the error file,err:%s", err)
			err = Errorerr
			return
		}

		f.Error = ErrorFile

	}

	return

}

func (f *LogFile) ValidateType() {
	SplitType := f.LogType

	switch SplitType {
	case LogWithHour:
		f.TimeType()
	case LogWithSize:
		f.SizeType()
	}

}

func (f *LogFile) writeLogInfo() {
	for message := range f.LogMsg {
		file := f.File
		if message.ErrorFatal {
			file = f.Error
		}
		//分割 日志

		f.ValidateType()

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
