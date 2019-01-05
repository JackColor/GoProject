/*
*Auth: JackColor
*Date: 2018/12/23 下午11:37
*/
package main

import "fmt"

func DemoDefer()  {

	i :=1
	defer fmt.Printf("first i %d\n",i)
	i=100
	fmt.Printf("last i %d\n",i)
}

func DemoDefer2()  {
	//defer 压栈 出栈
	defer fmt.Println("first")
	defer fmt.Println("second")


	fmt.Println("ooxx")




}




func main()  {

//DemoDefer()

DemoDefer2()




}


