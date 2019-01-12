/*
*Auth: JackColor
*Date: 2019/1/12 下午1:48
*/
package main

import "fmt"

//先 x 然后把 x赋值给 return 然后 defer x+1
func A() int {

	x := 5

	defer func() {

		x += 1

	}()

	return x

}

// return x = 5  defer x+1=6  show x=6
func B() (x int) { //返回值已经确定

	defer func() {
		x += 1
	}()

	return 5

}

// return x=y=5  defer x+1=6 show y =5
func C() (y int) {

	x := 5

	defer func() {
		x += 1
	}()

	return x

}

// return x=5  defer x=5副本 show x=5
func D() (x int) {

	defer func(x int) {
		x += 1

	}(x)

	return 5

}

// return x=5  defer 指针*x=6   show x=6
func E() (x int) {

	defer func(x *int) {
		*x += 1

	}(&x)

	return 5

}

func main() {

	//fmt.Println(A())
	//fmt.Println(B())
	//fmt.Println(C())
	fmt.Println(D())

}
