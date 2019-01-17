/*
*Auth: JackColor
*Date: 2019/1/14 下午9:27
*/
package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string `ini:"ip"`
	Port uint   `ini:"port"`
}

type MysqlConfig struct {
	Username string  `ini:"username"`
	Passwd   string  `ini:"passwd"`
	Database string  `ini:"database"`
	Host     string  `ini:"host"`
	Port     int     `ini:"port"`
	Timeout  float32 `ini:"timeout"`
}

func TestParser(t *testing.T) {

	data, err := ioutil.ReadFile("../config.ini")
	if err != nil {
		t.Error("failed read file name ")
	}
	var config Config
	//var  i int
	err = UnMarshal(data, &config)
	if err != nil {
		t.Error("UnMarshal failed info is ", err)
	}

	fmt.Printf("%#v", &config)

}
