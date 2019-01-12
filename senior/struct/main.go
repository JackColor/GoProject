/*
*Auth: JackColor
*Date: 2019/1/5 下午12:16
*/
package main

import (
	"fmt"
)

type User struct {
	Name   string
	Age    int
	Gender string
}

func Modify(u *User) {

	u.Name = "Modify"

}

func main() {

	var user User
	fmt.Printf("user %v", user)

	var user01 User = User{
		Name:   "alex",
		Age:    11,
		Gender: "女",
	}
	//还可以部分 赋值

	fmt.Printf("stu01 %v\n", user01)

	// 指针类型 初始化
	// 1
	var user02 *User = &User{}
	user02.Age = 13
	user02.Name = "Jack"
	fmt.Printf("stu02 %v\n", user02)
	//2
	var user03 *User = new(User)
	user03.Name = "user03"

	fmt.Printf("stu03 %v\n", user03)
	//3
	var user04 *User = &User{
		Name: "stu04",
	}
	fmt.Printf("user 04 %v\n", user04)

	Modify(user04)
	fmt.Printf("user 04 %v", user04)
}
