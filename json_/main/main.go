/*
*Auth: JackColor
*Date: 2019/1/14 下午8:33
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func main() {

	s := Student{
		Name:"alex",
		Age:11,
	}


	data,err:= json.Marshal(s)
	if err!=nil{
		fmt.Println(err)
	}

	fmt.Println(string(data))

	jsondata:= `{"name":"bob","age":11,"xxx":111}`

	var  s1 Student

	err2:= json.Unmarshal([]byte(jsondata),&s1)
	if err!=nil{
		fmt.Println(err2)
	}


	fmt.Println(s1.Name)
	fmt.Println(s1.Age)






}


