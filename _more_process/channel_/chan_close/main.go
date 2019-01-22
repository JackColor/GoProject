/*
*Auth: JackColor
*Date: 2019/1/17 下午10:44
*/
package main

import (
	"fmt"
	"time"
)

func Produce(c chan int) {

	for i := 0; i < 50; i++ {

		c <- i

	}
	//放进去关闭
	close(c)

	fmt.Println("closed!")
}

func main() {
	c := make(chan int, 100)
	go Produce(c)
	time.Sleep(time.Second * 5)
	//读取数据 直到读不到

	for {
		//一直获取 判断 chan 关闭用 ok 来表示
		data, ok := <-c
		if !ok {
			break
		}

		fmt.Println("get data ===>>", data)

	}

}
