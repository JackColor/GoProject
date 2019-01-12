/*
*Auth: JackColor
*Date: 2019/1/5 下午12:43
*/
package user

type User struct {
	Name   string
	Age    int
	Gender string
}

func GetUser(name string, age int, gender string) *User {

	return &User{
		Name: name,
		Age:  age,
		Gender: gender,
	}

}
