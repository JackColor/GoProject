/*
*Auth: JackColor
*Date: 2018/12/11 下午11:24
*/
package main

import "fmt"

func reverserOne(str string) {
	strbyte := []rune(str)
	length := len(strbyte)
	reverserStr := make([]rune, 0, 0)
	for index := range strbyte {

		reverserStr = append(reverserStr, strbyte[length-index-1])

	}
	fmt.Println(string(reverserStr))

}

func reverse(str string) {
	var res string
	length := len(str)
	for i := 0; i < length; i++ {

		res += fmt.Sprintf("%c", str[length-1-i])

	}
	fmt.Println(res)
}

func main() {

	reverserOne("gdgdgs")

}
