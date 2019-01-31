/*
*Auth: JackColor
*Date: 1/25/19 12:15 AM
*/
package main

import (
	"bufio"

	"fmt"
	"os"
	"strings"
	"study/Go-Ftp/ftp-protocol"
	"study/Go-Ftp/utils"
)

func (c *Client) Process() (err error) {

	//先 认证
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("please input the username:")
	Username, err := reader.ReadString('\n')
	fmt.Println("please input the password:")
	Password, err := reader.ReadString('\n')
	Username = strings.TrimSpace(Username)
	Password = strings.TrimSpace(Password)
	User := &ftp_protocol.User{
		Name:     Username,
		Password: Password,
		Status:   false,
	}

	user, status, err := c.Login(User)

	if !status {
		fmt.Println("login failed,username is ", user)
		return
	}
	c.User = User
	fmt.Println("user login successfully", user.Name)

	for {
		fmt.Println("please input the command:")
		Cmd, CommandError := reader.ReadString('\n')
		if CommandError != nil {
			fmt.Println("get the command failed", CommandError)
			return
		}

		Cmd = strings.TrimSpace(Cmd)
		fmt.Println(Cmd)
		switch Cmd {
		case ftp_protocol.ListDir:
			fmt.Println("list dir")
			c.ListDir()
		case ftp_protocol.SendFile:
			c.SendFile()
		case ftp_protocol.RecvFile:
		default:
			c.NotFind()

		}

	}

	return
}

func (c *Client) Login(UserInstance *ftp_protocol.User) (User *ftp_protocol.User, status bool, err error) {
	//认证

	if UserInstance.Password == "123" && UserInstance.Name == "alex" {
		UserInstance.Status = true
	}

	User = UserInstance
	User.Status = true
	status = true

	return
}

func (c *Client) ListDir() (err error) {
	fmt.Println("ls command start")
	currentPath := c.CurrentPath

	Msg := &ftp_protocol.Message{
		Cmd:    ftp_protocol.ListDir,
		Length: len(currentPath),

		Path: currentPath,
		User: c.User,
	}

	SendData, err := utils.Pack(Msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.Conn.Write(SendData)
	var data = make([]byte, 1024)
	dataLength := 0

	fmt.Println("here")

	for {

		MsgData := make([]byte, 0)

		n, errs := c.Conn.Read(data)
		if errs != nil {
			err = errs
			fmt.Println("read byte failed", err)
			return
		}
		dataStr := data[:n]
		//第一次接收
		c.Conn.Write([]byte(ftp_protocol.OK))

		if dataLength != 0 {
			MsgData = append(MsgData, dataStr...)

			fmt.Println(string(MsgData))

			allDateLength := len(MsgData)

			if allDateLength == dataLength {
				Message, _ := utils.UnPack(MsgData)
				for _, val := range Message.Info {
					fmt.Println(val)
				}
				break
			}
			continue
		}

		Msg, errs := utils.UnPack(dataStr)
		if errs != nil {
			err = errs
			return
		}

		if Msg.Info == nil {

			dataLength = Msg.Length
		}

	}

	return
}
func (c *Client) SendFile() (msg *ftp_protocol.Message, err error) {

	return
}

func (c *Client) NotFind() (msg *ftp_protocol.Message, err error) {

	return
}
