/*
*Auth: JackColor
*Date: 2019/1/6 下午2:37
*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFileSenior() {

	file, err := os.Open("./reader-senior/files/main.go")
	defer file.Close()
	if err != nil {
		fmt.Println("read fail failed", err)
		return
	}

	read := bufio.NewReader(file) //封装高级文件读写api

	for {

		line, err := read.ReadString('\n')

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file failed", err)
		}

		fmt.Println(line)

	}

}

func ReadFile() {

	file, err := os.Open("./reader-senior/files/main.go")

	if err != nil {
		fmt.Println("open file failed", err)

	}

	defer file.Close()
	var content []byte
	var buf [128]byte
	for {

		n, err := file.Read(buf[:])

		if err != nil {

			fmt.Println("read file failed", err)
		}

		if err == io.EOF {
			break
		}

		content = append(content, buf[:n]...)

		//fmt.Println("read content ", string(buf[:n]))
	}

	fmt.Println("read content ", string(content))

}

func main() {

	//reader := bufio.NewReader(os.Stdin)
	//
	//str,_:=reader.ReadString('\n')
	//
	//os.Stdout.WriteString(str)

	ReadFile()

}
