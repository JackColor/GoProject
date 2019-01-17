/*
*Auth: JackColor
*Date: 2019/1/14 下午7:25
*/
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age int
}

func modify_struct_value(a interface{})  {
	v:= reflect.ValueOf(a)

	v.Elem().Field(0).SetString("stu01")
	v.Elem().FieldByName("Age").SetInt(222)


}


func assert_struct(a interface{})  {

	v :=reflect.ValueOf(a) // 值
	t:= v.Type() // 类型
	fmt.Printf("the a is value is %v,and the type is %v\n",v,t)

	kind:= v.Kind()

	switch kind {

	case reflect.Int:
		fmt.Println("a is a int type")
	case reflect.Struct:
		fmt.Println("a is a struct")
		//fmt.Printf("the struct fields number is %d\n",v.NumField())
		for i:=0;i<v.NumField();i++{

			//获取每一个字段的的 类型的名称
			fmt.Printf("the struct's field's name is %v\n",t.Field(i).Name)
			//获取每一个字段的 type
			fmt.Printf("the struct's field's type is %v\n",v.Field(i).Type())
			// 值
			fmt.Printf("the struct's field's value is %v\n",v.Field(i).Interface())
		}

	default:
		fmt.Println("can't find the type kind")






	}












}





func main() {

	var s Student

	//assert_struct(s)
	modify_struct_value(&s)
	fmt.Println(s)
}


