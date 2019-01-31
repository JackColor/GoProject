/*
*Auth: JackColor
*Date: 1/23/19 12:40 AM
*/
package main

import (
	"fmt"
	"time"
)

func write(c chan int)  {

	for{
		select {
		case c<-123:
			fmt.Println("write date to channel")
		default:
			fmt.Println("chan is full wait....")


		}
		time.Sleep(time.Millisecond*50)
	}



}


func main() {

	chanInt :=make(chan int ,5)

	go write(chanInt)

	for v:=range chanInt{
		time.Sleep(time.Second)
		fmt.Println("get date from channel",v)


	}









}


