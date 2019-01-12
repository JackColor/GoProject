/*
*Auth: JackColor
*Date: 2019/1/12 下午5:26
*/
package main

import "fmt"

func empty() {
	var a interface{} //空接口可以接受任何值

	var b int = 100
	a = b

	fmt.Println("a value is ", a)

	var s string = "hello world"

	a = s

	fmt.Println("a value is ", a)

}

func seniorAssert(a interface{}) {

	switch v := a.(type) {

	case string:
		fmt.Printf("the a is string %s\n", v)
	case int:
		fmt.Printf("the a is int %d\n", v)

	case map[string]int:
		fmt.Printf("the a is a map[string]int %v\n", v)
	default:
		fmt.Printf("type not find from the assert %v\n", v)

	}

}

func assertType(a interface{}) {
	//断言 接口类型
	v, ok := a.(int)
	if ok {
		fmt.Printf("a is the int value is %d\n", v)
		return
	}

	s, ok := a.(string)
	if ok {
		fmt.Printf("a is the string value is %s\n", s)
		return
	}

	fmt.Printf("type not find from the assert %v", a)

}

func main() {

	var s string = "hello"
	//assertType(s)
	seniorAssert(s)
	var d int = 1
	//assertType(d)
	seniorAssert(d)

	m := make(map[string]int, 10)
	m["1"] = 1
	m["2"] = 2

	seniorAssert(m)

	f := 3.2
	seniorAssert(f)

}
