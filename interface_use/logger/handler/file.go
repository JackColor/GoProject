/*
*Auth: JackColor
*Date: 2019/1/12 下午7:47
 */
package handler

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type FileLog struct {
	path         string
	filename     string
	level        int
	file         *os.File
	errors       *os.File
	MsgChan      chan *LogData
	LogSplitType string
	LogSplitSize int64
	CurrentHour  int
}

func NewFileLog(config map[string]string) (log LogInterface, err error) {
	//返回 logInterface 的接口对象

	path, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("dont find the log path for this logger config map")
		return
	}

	filename, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("dont find the log name for this logger config map")
		return
	}

	logLevel, ok := config["level"]

	level := getLevel(logLevel)

	if !ok {
		err = fmt.Errorf("dont find the log level for this logger config map")
		return
	}

	MsgChanSize, ok := config["log_msg_size"]
	if !ok {
		MsgChanSize = "50000"

	}
	MsgSize, err := strconv.Atoi(MsgChanSize)
	if err != nil {
		MsgSize = 50000
	}

	splitType, ok := config["log_split_type"]

	if !ok {
		splitType = "hour"
	}

	var splitSize int64
	if splitType == "size" {

		splitSizeStr, ok := config["log_split_size"]
		if !ok {
			splitSizeStr = "104857600"
		}

		splitSize, err = strconv.ParseInt(splitSizeStr, 10, 64)
		if err != nil {
			splitSize = 104857600
		}

	}

	log = &FileLog{
		path:         path,
		filename:     filename,
		level:        level,
		MsgChan:      make(chan *LogData, MsgSize),
		CurrentHour:  time.Now().Hour(),
		LogSplitType: splitType,
		LogSplitSize: splitSize,
	}

	fmt.Println(log)
	log.Init()
	return

}

func (f *FileLog) Init() {
	// 创建 文件
	//初始化 init fileLog 对象

	filename := fmt.Sprintf("%s/%s.log", f.path, f.filename)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {

		panic(fmt.Sprintf("open filename [%s]is failed", filename))
	}

	f.file = file

	filename = fmt.Sprintf("%s/%s.log.wf", f.path, f.filename)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("open filename [%s]is failed", filename))
	}

	f.errors = file

	go f.writeLogData()

}

func (f *FileLog) SetLevel(level int) {
	// 设置 log 级别
	if level < Debug || level > Fatal {
		f.level = Info
		return
	}
	f.level = level
}

func (f *FileLog) Debug(format string, args ...interface{}) {
	if f.level > Debug {
		return
	}

	MsgData := writeLog(Debug, format, args...) // 把多个参数 展开传到函数里面  | 如果传入 args 这样就是一个参数
	select {
	case f.MsgChan <- MsgData:
	default:

	}

}

func (f *FileLog) Trace(format string, args ...interface{}) {
	if f.level > Trace {
		return
	}
	MsgData := writeLog(Trace, format, args...)
	select {
	case f.MsgChan <- MsgData:
	default:

	}
}

func (f *FileLog) Info(format string, args ...interface{}) {
	if f.level > Info {
		return
	}

	MsgData := writeLog(Info, format, args...)
	select {
	case f.MsgChan <- MsgData:
	default:

	}
}

func (f *FileLog) Warn(format string, args ...interface{}) {
	if f.level > Warn {
		return
	}

	MsgData := writeLog(Warn, format, args...)
	select {
	case f.MsgChan <- MsgData:
	default:

	}
}
func (f *FileLog) Error(format string, args ...interface{}) {
	if f.level > Error {
		return
	}
	MsgData := writeLog(Error, format, args...)
	select {
	case f.MsgChan <- MsgData:
	default:
	}

}
func (f *FileLog) Fatal(format string, args ...interface{}) {
	if f.level > Fatal {
		return
	}
	MsgData := writeLog(Fatal, format, args...)
	select {
	case f.MsgChan <- MsgData:
	default:

	}
}
func (f *FileLog) Close() {
	f.file.Close()
	f.errors.Close()

}

func (f *FileLog) splitWithHour(errorFatal bool) {
	// 时间分割
	now := time.Now()

	currentHour := now.Hour()

	if currentHour == f.CurrentHour {
		//不要备份....
		return
	}

	f.CurrentHour = currentHour
	var backupName string
	var oldFileName string
	if errorFatal {
		backupName = fmt.Sprintf("%s/%s.log.wf_%04d%02d%02%d02%d", f.path, f.filename, now.Year(), now.Month(), now.Day(), currentHour)
		oldFileName = fmt.Sprintf("%s/%s.log.wf", f.path, f.filename)
	} else {
		backupName = fmt.Sprintf("%s/%s.log_%04d%02d%02%d02%d", f.path, f.filename, now.Year(), now.Month(), now.Day(), currentHour)
		oldFileName = fmt.Sprintf("%s/%s.log", f.path, f.filename)
	}

	file := f.file
	if errorFatal {
		file = f.errors
	}
	file.Close()

	os.Rename(oldFileName, backupName)

	file, err := os.OpenFile(oldFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return
	}

	if errorFatal {
		f.errors = file
	} else {
		f.file = file
	}

}
func (f *FileLog) splitWithSize(errorFatal bool) {
	// 大小分割
	file := f.file
	if errorFatal {
		file = f.errors
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return
	}
	fileSize := fileInfo.Size()

	fmt.Println(fileSize)
	fmt.Println(f.LogSplitSize)

	if fileSize <= f.LogSplitSize {
		//小 返回
		fmt.Println("小 返回")
		return
	}

	fmt.Println("大")
	now := time.Now()
	var backupName string
	var oldFileName string
	if errorFatal {
		backupName = fmt.Sprintf("%s/%s.log.%04d%02d%02d%02d%02d", f.path, f.filename, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
		oldFileName = fmt.Sprintf("%s/%s.log.wf", f.path, f.filename)
	} else {
		backupName = fmt.Sprintf("%s/%s.log_%04d%02d%02d%02d%02d", f.path, f.filename, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())
		oldFileName = fmt.Sprintf("%s/%s.log", f.path, f.filename)
	}

	// -----///
	replaceFile := f.file
	if errorFatal {
		file = f.errors
	}
	replaceFile.Close()

	os.Rename(oldFileName, backupName)

	replaceFile, err = os.OpenFile(oldFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		return
	}

	if errorFatal {
		f.errors = replaceFile
	} else {
		f.file = replaceFile
	}

}

func (f *FileLog) checkSplitFlag(errorFatal bool) {
	// 判断什么切分方式

	if f.LogSplitType == SplitWithHour {

		f.splitWithHour(errorFatal)
	} else {

		f.splitWithSize(errorFatal)
	}

}

func (f *FileLog) writeLogData() {

	for data := range f.MsgChan {
		file := f.file
		if data.ErrorAndFatal {
			file = f.errors
		}

		f.checkSplitFlag(data.ErrorAndFatal)

		fmt.Fprintf(file, data.Msg)

	}

}
