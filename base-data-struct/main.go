/*
*Auth: JackColor
*Date: 2018/12/8 下午11:29
*/
package main

import "fmt"

func main() {

	m1 := make(map[string]interface{})

	m1["1"] = 1
	m1["2"] = "alex"
	age, ok := m1["1"].(int)
	if !ok {
		fmt.Println("not int") //断言
	}

	name, pk := m1["2"].(string)
	if !pk {
		fmt.Println("xxx")
	}

	fmt.Println(name)
	fmt.Println(age)

}
