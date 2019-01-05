/*
*Auth: JackColor
*Date: 2019/1/5 下午1:44
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
}

func main()  {
	stu := &Student{
		Name:"alex",
		Age:13,
		Gender:"女",
	}

	// 2进制
	data,_ := json.Marshal(stu)

	fmt.Printf("student %v",string(data))


	// 作业  //学生管理系统
}





