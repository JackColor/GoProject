/*
*Auth: JackColor
*Date: 2019/1/17 下午10:58
*/
package main

import (
	"fmt"
	"time"
)

func Produce(c chan int) {

	for i := 0; i < 50; i++ {

		c <- i
		time.Sleep(time.Millisecond)

	}
	//放进去关闭
	close(c)

	fmt.Println("closed!")
}

func main() {
	c := make(chan int)

	go Produce(c)

	for v := range c {
		//堵塞 直到chan 关闭
		fmt.Println("get data  ", v)

	}

}
