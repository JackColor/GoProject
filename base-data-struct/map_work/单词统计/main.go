/*
*Auth: JackColor
*Date: 2019/1/5 上午11:39
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func totalStr(line string) (charNum, DeistNum, SpaceNum, OtherNum int) {

	lineList := []rune(line)

	for _, v := range lineList {

		switch {
		case v <= 'z' && v >= 'a':
			charNum++
			fallthrough
		case v <= 'Z' && v >= 'A':
			charNum++
		case v == ' ':
			SpaceNum++
		case '0' <= v && v <= '9':
			DeistNum++
		default:
			OtherNum++

		}

	}
	return

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {

		return
	}

	char, gi, sp, oth := totalStr(string(line))

	fmt.Println(char, gi, sp, oth)

}
