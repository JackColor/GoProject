/*
*Auth: JackColor
*Date: 2019/1/5 下午1:50
*/
package main

import "fmt"

type Cat struct {
	Name   string
	Gender string
}

type Mint int

func (m *Mint) Print() {
	*m = 2000
	fmt.Printf("Mint is %d\n", *m)

}

func (c *Cat) Eat(f string) {
	fmt.Printf("%v is eacting %s\n", c.Name, f)
}

func main() {

	c := &Cat{
		Name:   "alex",
		Gender: "female",
	}

	c.Eat("fish")

	var num Mint = 100

	(&num).Print()

	//修改了
	fmt.Println(num)

}
