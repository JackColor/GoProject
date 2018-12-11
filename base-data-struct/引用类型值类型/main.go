/*
*Auth: JackColor
*Date: 2018/12/11 下午11:59
*/
package main

import "fmt"

func change(a, b *int) {

	temp := *a  // *a 取地址的值改变
	*a = *b
	*b = temp
}

func main() {

	a := 1
	b := 2

	a, b = b, a
	fmt.Println(a, b)
	change(&a, &b)
	fmt.Println(a, b)

}
