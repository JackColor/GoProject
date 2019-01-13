/*
*Auth: JackColor
*Date: 2019/1/12 下午8:15
*/
package app

import (
	"study/interface_use/logger/handler"
	"testing"
)

func TestLogger(t *testing.T) {
	///Users/mac/go/src/study/interface_use
	//log := handler.NewFileLog("/Users/mac/go/src/study/interface_use", "test")
	//
	//log.SetLevel(handler.Debug)
	//
	//log.Debug("test %s", "hello")
	//
	//log.Error("errors msg about is %s", "errors")

}

func TestConsoleLogger(t *testing.T) {

	log := handler.NewConsoleLog(handler.Debug)

	log.SetLevel(handler.Debug)

	log.Debug("test %s", "hello")

	log.Error("errors msg about is %s", "errors")

}
