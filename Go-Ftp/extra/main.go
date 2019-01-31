/*
*Auth: JackColor
*Date: 1/27/19 8:47 AM
*/
package main

import (
	"fmt"
	"io"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "www.baidu.com:80")

	if err != nil {
		fmt.Println("errors ", err)
		return
	}
	data := "GET / HTTP/1.1\r\n"
	data += "HOST: www.baidu.com\r\n"
	data += "connection:close\r\n"
	data += "\r\n\r\n"

	_, err = io.WriteString(conn, data)

	//conn.Write([]byte(data))

	if err != nil {
		return
	}

	var buf [1024]byte

	for {
		n, err := conn.Read(buf[:])
		if err != nil || n == 0 {
			break
		}

		fmt.Println(string(buf[:n]))
	}

}
