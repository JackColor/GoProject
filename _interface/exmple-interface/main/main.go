/*
*Auth: JackColor
*Date: 2019/1/12 下午5:04
*/
package main

import "fmt"

//定义接口 只要类实现全部接口的方法 则这个类就实现了这个接口
type Employee interface {
	Calc() float32
}

// 定义一种员工
type Programmer struct {
	Name     string
	Base     float32
	discount float32
}

// 构造方法
func getPro(name string, base, discount float32) *Programmer {

	return &Programmer{
		Name:     name,
		Base:     base,
		discount: discount,
	}

}

// 定义另外一种员工
type Sales struct {
	Name     string
	Base     float32
	discount float32
}

// 构造方法
func getSal(name string, base, discount float32) *Sales {

	return &Sales{
		Name:     name,
		Base:     base,
		discount: discount,
	}

}

// 一种员工 实现接口的 全部方法 这个接口有一个方法
func (p *Programmer) Calc() float32 {

	return p.Base

}

// 另一种员工 实现接口的 全部方法 这个接口有一个方法
func (p *Sales) Calc() float32 {

	return p.Base * p.discount

}

func CalcAllSalary(e []Employee) (count float32) {

	for _, val := range e {

		count += val.Calc()

	}

	return

}

func main() {
	var allEmployees []Employee
	p1 := getPro("p1", 3000, 0.1)
	p2 := getPro("p2", 8000, 0.1)

	s1 := getSal("s1", 2000, 0.4)
	s2 := getSal("s2", 6000, 0.8)

	allEmployees = append(allEmployees, p1)
	allEmployees = append(allEmployees, p2)
	allEmployees = append(allEmployees, s1)
	allEmployees = append(allEmployees, s2)

	count := CalcAllSalary(allEmployees)
	fmt.Println(count)

}
