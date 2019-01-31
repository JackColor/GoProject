/*
*Auth: JackColor
*Date: 1/25/19 12:14 AM
*/
package main

import (
	"fmt"
	"net"
	"study/Go-Ftp/ftp-protocol"
)

type Client struct {
	Conn        net.Conn
	Cmd         *ftp_protocol.Instruction
	CurrentPath string
	User        *ftp_protocol.User
}

func main() {

	conn, err := net.Dial("tcp", "0.0.0.0:8090")
	if err != nil {
		fmt.Println("client connect failed err:", err)

		return

	}

	fmt.Println("start the connect of :", conn)
	//把 conn 封装成 Client 类
	//等待用户的输入指令 等待输入指令 封装 发给 process
	ClientInstance := &Client{
		Conn:        conn,
		CurrentPath: ftp_protocol.CurrentPath,
	}

	ClientInstance.Process()

	//go ClientInstance.Process()

}
