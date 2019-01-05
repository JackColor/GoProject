/*
*Auth: JackColor
*Date: 2018/12/24 下午11:11
*/
package main

import "fmt"

func Test() {

	f1 := func(a, b int) int {

		return a + b

	}

	sum := f1(1, 2)

	fmt.Printf("sum= %d", sum)

}

func Demo() {

	var a int = 10
	// 匿名函数
	defer func() {
		fmt.Printf("defer a=%d\n", a)
	}()

	a = 1000

	fmt.Printf("a=%d\n", a)

}


//函数作为参数
func add(a,b int) int  {

	return a+b
}


func Calc(a,b int, functions func(int,int)int) int {

	return functions(a,b)



}

func Demo2()  {
	sum := Calc(100,200,add)

	fmt.Printf("sum=%d\n",sum)
}











func main() {

	Demo()
Demo2()
}
