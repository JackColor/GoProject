/*
*Auth: JackColor
*Date: 2019/1/5 下午7:17
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
}

type Class struct {
	Name     string
	Count    int
	Students []*Student
}

func main() {

	//序列化

	cls := &Class{
		Name: "101",
	}
	for i := 1; i <= 10; i++ {

		cls.Count++
		name := fmt.Sprintf("stu-%d", i)
		stu := &Student{
			Name: name,
		}
		cls.Students = append(cls.Students, stu)

	}

	data, err := json.Marshal(cls)
	if err != nil {
		fmt.Println("反序列化失败  line 48")

		return
	}

	fmt.Printf("data is %v", string(data))

	rowJson := `
{"Name":"101","Count":10,"Students":[{"Name":"stu-1"},{"Name":"stu-2"},{"Name":"stu-3"},{"Name":"stu-4"},{"Name":"stu-5"},{"Name":"stu-6"},{"Name":"stu-7"},{"Name":"stu-8"},{"Name":"stu-9"},{"Name":"stu-10"}]}

`
	cls2 := &Class{}

	err = json.Unmarshal([]byte(rowJson), cls2)
	if err != nil {
		fmt.Println("errors", err.Error())
		return
	}
	fmt.Println("check \n\n")
	fmt.Printf("name %s\n", cls2.Name)
	fmt.Printf("count %d\n", cls2.Count)
	for _, v := range cls2.Students {

		fmt.Printf("student is %s\n", v)

	}

}
