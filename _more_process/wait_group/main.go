/*
*Auth: JackColor
*Date: 2019/1/17 下午11:24
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Process(w *sync.WaitGroup)  {
	n:= rand.Intn(10)

	n2 := time.Duration(n)
	//fmt.Println(n2.String())
	fmt.Println("into the process")

	time.Sleep(n2)
	fmt.Println("dead the process")
	w.Done()



}


func main()  {
	var w sync.WaitGroup

	for i:=0;i<3;i++{


		w.Add(1)
		go Process(&w)


	}


	w.Wait()

	fmt.Println("main func is dead")









}


