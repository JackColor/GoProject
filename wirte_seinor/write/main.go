/*
*Auth: JackColor
*Date: 2019/1/12 下午12:31
*/
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writer_utils() {

	var s string = "hello world"

	e := ioutil.WriteFile("./test.dat", []byte(s), 0777)

	if e != nil {

		return
	}

}

func Writer_buffer() {

	file, err := os.OpenFile("./test.dat", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("open file failed", err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {

		writer.WriteString("hello world\n")

	}

	writer.Flush()

}

func Wirter() {

	file, err := os.OpenFile("./test.dat", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("open file failed", err)
	}

	defer file.Close()

	var s string = "hello world\n"

	n, err := file.Write([]byte(s))

	file.WriteString(s)
	if err != nil {
		fmt.Println("write content failed", err)
	}
	fmt.Println(n)

}

func main() {
	//Writer_buffer()
	writer_utils()
}
