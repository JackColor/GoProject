/*
*Auth: JackColor
*Date: 2019/1/17 下午8:27
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {

	cpu := runtime.NumCPU()

	runtime.GOMAXPROCS(2)

	fmt.Println("this computer is ", cpu)

}
