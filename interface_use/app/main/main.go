/*
*Auth: JackColor
*Date: 2019/1/12 下午11:00
*/
package main

import (
	"fmt"
	"study/interface_use/logger/handler"
	"time"
)

func main() {

	m1 := make(map[string]string)
	m1["log_path"] = "/Users/mac/go/src/study/interface_use"
	m1["log_name"] = "test2"
	m1["level"] = "Debug"
	m1["log_split_type"] = "size"
	m1["log_split_size"] = "10000"
	err := handler.InitLogger("file", m1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		handler.DebugLog("this is the debug logger ")
		handler.WarnLog("this is the warn log")
		time.Sleep(time.Second)
	}

}
