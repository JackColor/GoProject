/*
*Auth: JackColor
*Date: 2019/1/1 下午1:18
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func work(s []string) {
	fmt.Println(s)
	m1 := make(map[string]int, 12)
	for _, val := range s {

		_, ok := m1[val]
		if !ok {
			m1[val] = 1
		} else {
			m1[val]++
		}
	}
	fmt.Printf("the result is %v", m1)
}

func main() {
	//读取屏幕输入
	reader := bufio.NewReader(os.Stdin)
	inputStr, err := reader.ReadString('a')
	if err != nil {

		return
	}
	fmt.Println(inputStr)
	//str := string(inputStr)

	strSlice := strings.Split(inputStr, " ")

	if len(strSlice) == 0 {
		return
	}

	work(strSlice)

}
