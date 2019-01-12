/*
*Auth: JackColor
*Date: 2019/1/5 下午1:38
*/
package main

import "fmt"

type City struct {
	Address string
	Phone   string
}

type User struct {
	Name string
	Age  int
	*City
}

func main() {
	// 1
	user := &User{
		Name: "alex",
		Age:  12,
		City: new(City),
	}
	user.Address = "北京"
	user.Phone = "12345"
	fmt.Printf("user ->%v\n", user)

	//2
	user2 := &User{}

	user2.City = &City{}

	user2.Address = "上海"
	user2.Name = "Jack"
	user2.Phone = "12321"

	fmt.Printf("user ->%v", user2)

}
