/*
*Auth: JackColor
*Date: 1/24/19 12:22 AM
*/
package main

import (
	"bufio"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8090")
	if err != nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)

	for {

		str, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		conn.Write([]byte(str))
	
	}




}
