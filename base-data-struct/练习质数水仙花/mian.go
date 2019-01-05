/*
*Auth: JackColor
*Date: 2018/12/24 上午12:04
*/
package main

import "fmt"

func justify(n int) bool {

	for i := 2; i < n; i++ {

		if n%i == 0 {
			return false
		}
	}
	return true
}

func Demo1() {

	for i := 2; i < 100; i++ {
		if justify(i) {

			fmt.Printf("%d is prima\n", i)

		}

	}

}

func isDaffodils(i int) bool {

	first := i%10
	second := (i/10)%10

	thrid :=(i/100)%10

	sum := first*first*first + second*second*second + thrid*thrid*thrid

	return i == sum

}


func Demo2() {

	//水仙花数
	for i:=100;i<1000;i++{

		if isDaffodils(i) {

			fmt.Printf("%d is daffodils\n",i)

		}



	}






}

func main() {

	//Demo1()
	Demo2()

}
