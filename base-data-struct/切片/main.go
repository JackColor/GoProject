/*
*Auth: JackColor
*Date: 2018/12/23 上午11:22
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Demo() {

	var a [5]int
	a[1] = 1

	var b [6]int = [6]int{1, 2, 3, 4}

	fmt.Printf("b is the ===>%d\n", b)

	c := []int{1, 2, 3, 4}

	fmt.Println(c)

	d := [4]int{2: 200, 3: 300}

	fmt.Printf("d is ===>%d", d)

}

func Demo2() {

	var b [6]int = [6]int{1, 2, 3, 4}

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	for index, val := range b {
		fmt.Println(index)
		fmt.Println(val)

	}

}

func getSum(a []int) (sum int) {
	//var sum int = 0
	for _, val := range a {

		sum += val

	}

	return

}

func testSum() {

	rand.Seed(time.Now().Unix())

	var testArray [20]int

	for i := 0; i < len(testArray); i++ {

		testArray[i] = rand.Intn(999)

	}

	res := getSum(testArray[:])

	fmt.Println(res)

}

func Demo3() {

	nums := []int{1, 3, 5, 7, 1}

	for index, val := range nums {

		for in, value := range nums {

			if (val + value) == 8 {

				println(index, in)

			}

		}

	}

}

func Demo4() {

	// 创建切片

	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	//

	s1 := a[1:]

	fmt.Println(s1)

	s2 := make([]int, 5, 5)

	fmt.Println(s2)

}

func main() {

	Demo4()
}
