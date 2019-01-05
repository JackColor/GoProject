/*
*Auth: JackColor
*Date: 2018/12/23 下午11:32
*/
package main

import "fmt"

func calcOne(a,b int) (sum int){

	sum = a+b


	return

}

func calcMore(a int, b ...int) (sum int) {

	sum = a



	for _,val :=range b{

		sum+=val


	}
	return

}





func main()  {

	s := calcMore(1,1,2,3,4,5,)
	fmt.Println(s)


}


