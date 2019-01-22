/*
*Auth: JackColor
*Date: 2019/1/17 下午9:05
*/
package main

import "fmt"

func hello(c chan bool) {
	c <- true

	fmt.Println("hello world")

}

func main() {
	c := make(chan bool)

	go hello(c)

	<-c
	fmt.Println("main func dead")

}
