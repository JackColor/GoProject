/*
*Auth: JackColor
*Date: 1/25/19 12:17 AM
*/
package ftp_protocol

type Instruction struct {
	Cmd string
}

type Message struct {
	Cmd    string  `json:"cmd"`
	Length int		`json:"length"`
	Info   []string    `json:"info"`
	Path string `json:"path"`
	User   *User	`json:"user"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   bool	`json:"status"`
}

const (
	ListDir  = "ls"
	SendFile = "put"
	RecvFile = "get"
	CurrentPath = "./"
	OK = "ok"
)
