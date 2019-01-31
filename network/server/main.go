/*
*Auth: JackColor
*Date: 1/23/19 11:59 PM
*/
package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()


	for {
		content := make([]byte, 1024)

		n, err := conn.Read(content)

		if err != nil {
			return
		}
		fmt.Printf("get date of %s", string(content[:n+1]))
	}

}

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:8090")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {

		conn, err := listen.Accept()

		if err != nil {
			return
		}

		fmt.Println("conn", conn)

		go process(conn)
	}
}
