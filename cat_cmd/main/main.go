/*
*Auth: JackColor
*Date: 2019/1/12 下午1:04
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var n int

func cat(reader *bufio.Reader) {
	count := 0
	for {

		s, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		count++
		if count == n {
			break
		}
		fmt.Println(s)
	}

}

func init() {

	flag.IntVar(&n, "n", 10, "n is the hang shu")

	flag.Parse()

	fmt.Println("n is ", n)

}
func main() {

	//
	if flag.NArg() == 0 {
		reader := bufio.NewReader(os.Stdin)
		cat(reader)

	}

	if flag.NArg() > 0 {

		for _, val := range flag.Args() {

			file, err := os.Open(val)
			if err != nil {

				fmt.Println("open file failed", err)
			}
			reader := bufio.NewReader(file)
			cat(reader)
		}

	}

}
