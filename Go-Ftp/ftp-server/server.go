/*
*Auth: JackColor
*Date: 1/26/19 7:27 PM
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"study/Go-Ftp/ftp-protocol"
	"study/Go-Ftp/utils"
)

func (s *Server) Process() (err error) {

	for {
		Msg := make([]byte, 1024)
		n, errs := s.Conn.Read(Msg)
		if errs != nil {
			fmt.Println("server process err:", errs)
			err = errs
			return
		}

		Message, errs := utils.UnPack(Msg[:n])
		if errs == io.EOF {

			return
		}

		if err != nil {
			fmt.Println("UuMarshal failed err:", err)
			return
		}

		fmt.Println(Message)

		s.Msg = Message

		switch Message.Cmd {
		case ftp_protocol.ListDir:
			s.LisDir()
		case ftp_protocol.SendFile:
			s.SendFile()
		case ftp_protocol.RecvFile:
			s.RecvFile()

		}

	}

	return
}

func (s *Server) LisDir() (err error) {

	files, err := ioutil.ReadDir(s.CurrentPath)

	if err != nil {
		fmt.Println("read dirs failed err:", err)
	}

	SendStringSlice := make([]string, 0)
	for _, v := range files {
		NowStr := v.Name()
		Flag := strings.Index(NowStr, ".")
		if Flag == -1 {

			SendStringSlice = append(SendStringSlice, NowStr)
		}

	}

	MsgData := &ftp_protocol.Message{
		Cmd:  s.Msg.Cmd,
		Info: SendStringSlice,
		User: s.User,
	}

	msgdata, err := utils.Pack(MsgData)

	dataLength := len(msgdata)

	if err != nil {
		fmt.Println("marshal failed in files info err:", err)
	}

	if dataLength > 2046 {

		//循环发送

	}

	if err != nil {

		fmt.Println("msgdata failed", err)
		return
	}

	Msg := &ftp_protocol.Message{
		Cmd:    s.Msg.Cmd,
		Length: len(msgdata),
		User:   s.User,
	}

	data, err := utils.Pack(Msg)

	if err != nil {
		return
	}

	fmt.Println("....footer")

	s.Conn.Write(data)

	okDate := make([]byte, 1024)

	n, _ := s.Conn.Read(okDate)

	okFlag := string(okDate[:n])

	fmt.Println("OK", okFlag)

	if okFlag != ftp_protocol.OK {
		return
	}

	s.Conn.Write(msgdata)

	return
}

func (s *Server) SendFile() (err error) {





	return
}

func (s *Server) RecvFile() (err error) {

	return
}
