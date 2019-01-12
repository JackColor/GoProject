/*
*Auth: JackColor
*Date: 2019/1/6 下午2:50
*/
package main

import (
	"flag"
	"fmt"
)

var (
	index bool
	age   int
	name  string
)

func init() {
	flag.BoolVar(&index, "b", false, "bool args")
	flag.StringVar(&name, "n", "alex", "string args")
	flag.IntVar(&age, "i", 1, "int args")

	flag.Parse()

}

func main() {
	fmt.Printf("bool is %v\n", index)
	fmt.Printf("string is %s\n", name)
	fmt.Printf("int is %d\n", age)

}
