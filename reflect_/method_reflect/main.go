/*
*Auth: JackColor
*Date: 2019/1/14 下午8:16
*/
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s *Student) SetName(name string) {
	fmt.Println("this method is set name ")
	s.Name = name
}

func (s *Student) GetName() {
	fmt.Println("this method is get name")
	fmt.Println(s.Name)
}

func get_struct_name(a interface{}) {
	v := reflect.ValueOf(a)

	t := v.Type()

	num := t.NumMethod()
	//获取结构体的方法  注意指针类型与 值类型
	fmt.Println("the struct method's number is ", num)

	for i := 0; i < num; i++ {
		//函数的名字
		fmt.Printf("the method's name is %v\n", v.Type().Method(i).Name)
		// 函数的 形式 入参 出餐
		fmt.Printf("the method's sturct is %v\n", v.Method(i).Type())

	}
	//用 value的返回的值来 获得函数的签名
	m := v.MethodByName("GetName")
	var arg []reflect.Value

	m.Call(arg)

	m2 := v.MethodByName("SetName")
	var args2 []reflect.Value // 参数切片
	name := "stu01"
	// 参数转换
	nameVal := reflect.ValueOf(name)

	args2 = append(args2, nameVal)

	m2.Call(args2)

	m.Call(arg)

}

func main() {

	var s Student

	get_struct_name(&s)

}
