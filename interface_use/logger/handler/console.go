/*
*Auth: JackColor
*Date: 2019/1/12 下午9:59
*/
package handler

import (
	"fmt"
	"os"
)

type ConsoleLog struct {
	level int
}

func NewConsoleLog(config map[string]string) (logger LogInterface, err error) {

	logLevel, ok := config["level"]

	if !ok {
		err = fmt.Errorf("console log cant find the log level ")
		return
	}
	level := getLevel(logLevel)
	logger = &ConsoleLog{
		level: level,
	}

	return

}

func (c *ConsoleLog) Init() {

}

func (c *ConsoleLog) SetLevel(level int) {
	if level < Debug || level > Fatal {
		c.level = Debug
	}

	c.level = level

}

func (c *ConsoleLog) writeLog(file *os.File, level int, format string, args ...interface{}) {
	if c.level > level {
		return
	}

}

func (c *ConsoleLog) Debug(format string, args ...interface{}) {

	if c.level > Debug {
		return
	}
	MsgData := writeLog(Debug, format, args...)

	fmt.Fprintf(os.Stdout, MsgData.Msg)

}

func (c *ConsoleLog) Trace(format string, args ...interface{}) {
	if c.level > Trace {
		return
	}
	MsgData := writeLog(Trace, format, args...)
	fmt.Fprintf(os.Stdout, MsgData.Msg)
}

func (c *ConsoleLog) Info(format string, args ...interface{}) {
	if c.level > Info {
		return
	}
	MsgData := writeLog(Info, format, args...)
	fmt.Fprintf(os.Stdout, MsgData.Msg)
}

func (c *ConsoleLog) Warn(format string, args ...interface{}) {
	if c.level > Warn {
		return
	}
	MsgData := writeLog(Warn, format, args...)
	fmt.Fprintf(os.Stdout, MsgData.Msg)
}
func (c *ConsoleLog) Error(format string, args ...interface{}) {
	if c.level > Error {
		return
	}
	MsgData := writeLog(Error, format, args...)
	fmt.Fprintf(os.Stdout, MsgData.Msg)
}
func (c *ConsoleLog) Fatal(format string, args ...interface{}) {
	if c.level > Fatal {
		return
	}
	MsgData := writeLog(Fatal, format, args...)
	fmt.Fprintf(os.Stdout, MsgData.Msg)
}
func (c *ConsoleLog) Close() {

}
