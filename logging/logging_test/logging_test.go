/*
*Auth: JackColor
*Date: 2019/1/13 下午4:47
*/
package logging_test

import (
	"fmt"
	"study/logging"
	"testing"
	"time"
)

func TestFileLog(t *testing.T) {
	conf := make(map[string]string, 10)
	conf["log_path"] = "/Users/mac/go/src/study/logging/log_store"
	conf["log_name"] = "demo"
	conf["level"] = "DEBUG"
	conf["log_type"] = "size"
	conf["log_size"] = "1000"
	err := logging.Init("file", conf)
	if err != nil {
		fmt.Println(err)
	}
	for {
		logging.Debug("this is the info log %s", "info")
		logging.Error("this is the errors log %s", "errors")
		time.Sleep(time.Second)
	}

}
