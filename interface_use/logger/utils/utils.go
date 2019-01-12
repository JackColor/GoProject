/*
*Auth: JackColor
*Date: 2019/1/12 下午8:58
*/
package utils

import (
	"fmt"
	"runtime"
)

func GetLogLine() (FileName, FuncName string, Line int) {

	pc, filename, line, ok := runtime.Caller(0)

	if ok {
		FileName = filename
		FuncName = runtime.FuncForPC(pc).Name()
		Line = line
		return
	}
	fmt.Println("get log line info errors")
	return
}
