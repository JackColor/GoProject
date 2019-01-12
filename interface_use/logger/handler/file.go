/*
*Auth: JackColor
*Date: 2019/1/12 下午7:47
*/
package handler

import (
	"fmt"
	"os"
	"path"
	"study/interface_use/logger/utils"
	"time"
)

type FileLog struct {
	path     string
	filename string
	level    int
	file     *os.File
	errors   *os.File
}

// 返回接口
func NewFileLog(path, filename string) logInterface {
	//返回 logInterface 的接口对象
	logger := &FileLog{
		path:     path,
		filename: filename,
		level:    Debug,
	}

	logger.Init()
	return logger

}

func (f *FileLog) Init() {
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
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	filename, funcname, line := utils.GetLogLine()
	filename = path.Base(filename)
	funcname = path.Base(funcname)
	level := getLevelText(f.level)

	// 时间 级别 文件名称 函数名 函数行号  格式字符串
	msg := fmt.Sprintf("%s-%s %s %s:%d %s \n", nowStr, level, filename, funcname, line, format)
	//fmt.Fprintf(f.file, msg)

	fmt.Fprintf(f.file, msg, args)

}

func (f *FileLog) Trace(format string, args ...interface{}) {
	if f.level > Trace {
		return
	}
	fmt.Fprintf(f.file, format, args)
}

func (f *FileLog) Info(format string, args ...interface{}) {
	if f.level > Info {

		return
	}

	fmt.Fprintf(f.file, format, args)
}

func (f *FileLog) Warn(format string, args ...interface{}) {
	if f.level > Warn {
		return
	}

	fmt.Fprintf(f.file, format, args)
}
func (f *FileLog) Error(format string, args ...interface{}) {
	if f.level > Error {
		return

	}

	fmt.Fprintf(f.errors, format, args)
}
func (f *FileLog) Fatal(format string, args ...interface{}) {
	if f.level > Fatal {
		return

	}

	fmt.Fprintf(f.errors, format, args)
}
func (f *FileLog) Close() {
	f.Close()
}
