/*
*Auth: JackColor
*Date: 2019/1/12 下午12:48
*/
package main

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(dst, src string) (n int64, err error) {

	filesrc, err := os.Open(src)
	if err != nil {

		fmt.Println("open file failed ", err)
	}

	filedst, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {

		fmt.Println("write file content failed", err)
	}

	defer filesrc.Close()
	defer filedst.Close()

	n, err = io.Copy(filedst, filesrc)

	return

}

func main() {
	CopyFile("test3.dat","test.dat")



}
