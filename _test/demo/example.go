/*
*Auth: JackColor
*Date: 2019/1/16 下午11:46
*/
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Add(a, b int) int {

	return a + b

}

func Sub(a, b int) int {

	return a - b
}

var numChan chan int

func makeNums()  {
	numChan = make(chan int ,300)
	s := []int{3, 3, 3, 5, 5, 5, 5, 5, 2, 2}

	for {
		var randomNum int
		rand.Seed(time.Now().UnixNano())
		num := rand.Intn(len(s))
		r := s[num]
		switch r {
		case 3:
			randomNum = rand.Intn(8) +18 //[0,n) [18,n) [18,25]
		case 5:
			randomNum = rand.Intn(10) + 26//[0,n),==>(26, 35)
		case 2:
			randomNum = rand.Intn(65) + 36//[0,n),==>(36, 100)
		}
		//time.Sleep(time.Millisecond)
		//fmt.Println(randomNum)
		numChan<-randomNum

	}




}

func toKafka(w *sync.WaitGroup)  {

	for v :=range numChan{

		fmt.Println("get nums",v)


		/*


		es
		go





		*/
	}


 w.Done()
}






func main() {
	/*

	 3: (18, 25),
        5: (26, 35),
        2: (36, 100)
	*/
	w :=sync.WaitGroup{}
	for i:=0;i<100;i++{
		w.Add(1)
		go toKafka(&w)
	}
	 makeNums()
	//go makeNums()
	//go toKafka()



	w.Wait()




}
