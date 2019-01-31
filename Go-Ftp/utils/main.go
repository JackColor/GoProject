/*
*Auth: JackColor
*Date: 1/27/19 9:04 AM
*/
package utils

import (
	"encoding/json"
	"fmt"
	"study/Go-Ftp/ftp-protocol"
)

func Pack(msg *ftp_protocol.Message) (info []byte, err error) {

	info, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("marshal failed err:", err)
		return
	}

	return

}

func UnPack(msg []byte) (Msg *ftp_protocol.Message, err error) {

	Msg = &ftp_protocol.Message{}

	err = json.Unmarshal(msg, Msg)

	if err != nil {
		//fmt.Println(err)
		return
	}

	return
}
