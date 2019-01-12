/*
*Auth: JackColor
*Date: 2019/1/12 下午7:03
*/
package main

import (
	"fmt"
	"os"
	"study/interface_use/example/log"
)

func main() {
	file, err := os.OpenFile("./test.data", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed", err)
	}

	defer file.Close()

	//logHandler := log.NewFileLogHandler("./c/text.log")
	logHandler := log.NewConsoleLogHandler("./c/text.log")
	logHandler.LogDebug()
	logHandler.LogWarn()

}
