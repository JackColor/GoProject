/*
*Auth: JackColor
*Date: 1/22/19 11:34 PM
*/
package main

import (
	"fmt"
	"time"
)

func server(c chan string) {

	time.Sleep(time.Second * 3)
	c <- "server s1 input \n"

}

func server2(c chan string) {
	time.Sleep(time.Second * 5)
	c <- "server s2 input \n"

}

func main() {

	c := make(chan string, 2)

	go server(c)
	go server2(c)
	//time.Sleep(time.Second*4)
	select {
	case s1 := <-c:
		fmt.Println("get the data from s1", s1)
	case s2 := <-c:
		fmt.Println("get the data from s2", s2)
		//default:
		//	fmt.Println("no date ")
	}

}
