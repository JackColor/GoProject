/*
*Auth: JackColor
*Date: 2019/1/12 下午12:09
*/
package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./reader-senior/files/main.go.gz")

	if err != nil {

		fmt.Println("open file failed", err)
	}

	reader, err := gzip.NewReader(file)

	defer reader.Close()

	if err != nil {
		fmt.Println("get reader failed ", err)
	}

	newReader := bufio.NewReader(reader)

	for {

		line, err := newReader.ReadString('\n')
		if err == io.EOF {

			break
		}
		if err != nil {

			fmt.Println("read file failed", err)
		}

		fmt.Println(line)

	}

}
