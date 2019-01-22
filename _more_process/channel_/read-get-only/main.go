/*
*Auth: JackColor
*Date: 2019/1/17 下午10:31
*/
package main

import (
	"fmt"
)

func ReadOnly(c <-chan int) {

	data := <-c
	fmt.Println("read from chan ", data)

}

func SendOnly(c chan<- int) {

	a := 10

	c <- a

}

func main() {
	c := make(chan int)

	go SendOnly(c)

	ReadOnly(c)

	fmt.Println("main is over")
}
