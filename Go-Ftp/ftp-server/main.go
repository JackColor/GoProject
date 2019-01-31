/*
*Auth: JackColor
*Date: 1/27/19 9:22 AM
*/
package main



import (
	"fmt"
	"net"
	"study/Go-Ftp/ftp-protocol"
)


type Server struct {
	Conn        net.Conn
	Msg         *ftp_protocol.Message
	User        *ftp_protocol.User
	Cmd         string
	CurrentPath string
}
func main() {

	server, err := net.Listen("tcp", "0.0.0.0:8090")
	fmt.Println("start server in 8090")
	if err != nil {

		fmt.Println("server listen errors:", err)
		return

	}
	for {

		conn, err := server.Accept()
		if err != nil {
			fmt.Println("get the connections failed", err)
			return
		}

		fmt.Println("accept a conn ",conn)

		ServerInstance := &Server{
			Conn:conn,
			CurrentPath:ftp_protocol.CurrentPath,
		}

		go ServerInstance.Process()





	}

}