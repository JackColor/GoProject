/*
*Auth: JackColor
*Date: 1/23/19 11:07 PM
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

var RwLock = sync.RWMutex{}// 读写锁
var wg sync.WaitGroup
var x int = 1

func write() {

	RwLock.Lock() //拿到锁 之后 所有的锁等待释放
	x++
	time.Sleep(time.Second * 5)
	RwLock.Unlock()
	wg.Done()
}

func read(i int) {

	RwLock.RLock()
	fmt.Printf("go %d read the x is %d\n", i, x)
	RwLock.RUnlock()
	wg.Done()

}

func main() {

	wg.Add(1)
	go write()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read(i)

	}

	wg.Wait()

}
