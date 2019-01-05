/*
*Auth: JackColor
*Date: 2018/12/29 下午8:48
*/
package main

import (
	"fmt"
	"time"
)

func Print(i int){


	fmt.Println("this is ==>>",i)



}

func main() {

	for i := 1; i < 1000000; i++ {

		go Print(i)

	}


	time.Sleep(time.Second*1)


}
