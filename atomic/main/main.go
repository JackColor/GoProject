/*
*Auth: JackColor
*Date: 1/23/19 11:24 PM
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int32 = 1
var wg sync.WaitGroup

func add() {

	atomic.AddInt32(&x, 2)

	wg.Done()
}

func main() {

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add()

	}

	wg.Wait()

	fmt.Println("date x is", x)
}
