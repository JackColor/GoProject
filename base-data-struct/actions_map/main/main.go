/*
*Auth: JackColor
*Date: 2019/1/1 下午12:56
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Testmap() {

	m1 := make(map[string]int, 128)

	for i := 0; i < 10; i++ {

		key := fmt.Sprintf("stu%d", i+1)
		val := rand.Intn(1000)

		m1[key] = val

	}

	//fmt.Printf("the map is contains :%v",m1)
	delete(m1,"stu3")
	for key, val := range m1 {

		fmt.Printf("the key is: %s", key)
		fmt.Printf("the value is: %d\n", val)

	}







}

func main() {
	rand.Seed(time.Now().UnixNano())

	Testmap()

}
