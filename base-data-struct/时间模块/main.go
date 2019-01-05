/*
*Auth: JackColor
*Date: 2018/12/12 上午12:03
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	//获取当前时间戳
	now := time.Now()
	fmt.Println(now)
	time.Sleep(time.Millisecond * 1000)


	 fmt.Println(now.Format("2006=01=02 15:04:05"))



	//end := time.Now()
	//
	//fmt.Println("%v", end.Sub(now))
	//
	//h, _ := time.ParseDuration("4h30m") //时间段的分解
	//fmt.Printf("I've got %.1f mins of work left.", h.Minutes())

}
