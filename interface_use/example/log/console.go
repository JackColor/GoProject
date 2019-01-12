/*
*Auth: JackColor
*Date: 2019/1/12 下午6:58
*/
package log

import "fmt"

type ConsoleLog struct {
	path string
}

func NewConsoleLogHandler(path string) logInterface {

	return &ConsoleLog{
		path: path,
	}

}

func (f *ConsoleLog) LogWarn() {

	fmt.Printf("print warn log in to the console %s\n", f.path)
}

func (f *ConsoleLog) LogDebug() {
	fmt.Printf("print debug log in to the console %s\n", f.path)
}
