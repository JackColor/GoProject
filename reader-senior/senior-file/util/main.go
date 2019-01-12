/*
*Auth: JackColor
*Date: 2019/1/12 上午11:59
*/
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	bf,err:= ioutil.ReadFile("./reader-senior/files/main.go")
	if err!=nil{
		fmt.Println("read file failed",err)
	}
	fmt.Println(string(bf))

}




