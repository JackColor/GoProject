/*
*Auth: JackColor
*Date: 2019/1/12 下午3:47
*/
package main

import "fmt"

type Animal interface {
	Eat()
	Talk()
	Name() string
}

type Dog struct {
	name string
}

func (d *Dog) Eat() { //指针类型的 dog 实现了 animal 的方法

	fmt.Println("dog is eating")

}

func (d *Dog) Talk() {

	fmt.Println("dog is talk")
}

func (d *Dog) Name() string {

	fmt.Println("dog is name ", d.name)

	return d.name

}

func main() {

	var d *Dog = &Dog{name: "旺财"}
	var a Animal
	a = d
	a.Eat()
	a.Talk()
	a.Name()

}
