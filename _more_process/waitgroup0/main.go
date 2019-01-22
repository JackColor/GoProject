/*
*Auth: JackColor
*Date: 1/22/19 9:36 PM
*/
package main

import (
	"fmt"
	"time"
)

func Say(c chan bool)  {

	fmt.Println("hello--->")
	time.Sleep(time.Second*2)
	c<-true



}



func main() {

	no:=3

	flag:=make(chan bool)

	for i:=0;i<no;i++{

		go Say(flag)


	}


	for i:=0;i<no;i++{

		<-flag
	}






}


