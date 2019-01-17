/*
*Auth: JackColor
*Date: 2019/1/13 下午6:40
*/
package main

import (
	"fmt"
	"reflect"
)

func reflect_type(a interface{}) {

	v := reflect.TypeOf(a)
	fmt.Println("the a type is ", v)

	k := v.Kind()

	switch k {
	case reflect.Int:
		fmt.Println("a is the: ", reflect.Int)
	case reflect.Float32:
		fmt.Println("a is the: ", reflect.Float32)

	}

}

func reflect_value(a interface{}) {
	v := reflect.ValueOf(a)
	fmt.Println("the a value is ", v)

	k := v.Kind()

	switch k {

	case reflect.Int:
		fmt.Println("the a is the type is ", reflect.Int)
	case reflect.Float32:
		fmt.Println("the a is the type is ", reflect.Float32)

	case reflect.String:
		fmt.Println("the a is the type is ", reflect.String)

	}

}

func reflect_modify(a interface{}) {

	v := reflect.ValueOf(a)
	fmt.Println("the a value is ", v)

	k := v.Kind()

	switch k {

	case reflect.Int:
		fmt.Println("the a is the type is ", reflect.Int)
	case reflect.Float32:
		fmt.Println("the a is the type is ", reflect.Float32)

	case reflect.String:
		fmt.Println("the a is the type is ", reflect.String)
	case reflect.Ptr:

		fmt.Println("the a is the type is", reflect.Ptr)
		// 指针设置值
		v.Elem().SetFloat(32.2)

	}

}
func main() {
	var a float64 = 23.2

	//reflect_type(a)
	//reflect_value(a)
	reflect_modify(&a)
	fmt.Println(a)
}
