/*
*Auth: JackColor
*Date: 2019/1/17 下午8:20
*/
package main

import (
	"fmt"
	"time"
)

func get_number() {

	for i := 0; i <= 5; i++ {
		time.Sleep(time.Millisecond * 200)
		fmt.Printf("%d  ", i)
	}

}

func get_charaters() {

	for i := 'a'; i < 'h'; i++ {
		time.Sleep(time.Millisecond * 300)
		fmt.Printf("%c  ", i)

	}

}

func main() {

	go get_charaters()
	go get_number()
	time.Sleep(time.Second*10)
	fmt.Println("the main is dead")

}
