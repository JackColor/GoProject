/*
*Auth: JackColor
*Date: 2019/1/6 下午2:17
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	//var buf [16]byte
	////输入
	//os.Stdin.Read(buf[:])
	////输出
	//os.Stdout.WriteString(string(buf[:]))

	var a string
	var b int

	//fmt.Fscanf(os.Stdin, "%s%d\n", &a, &b)

	fmt.Fscan(os.Stdin, &a, &b)
	fmt.Fprintln(os.Stdout, a, b)

}
