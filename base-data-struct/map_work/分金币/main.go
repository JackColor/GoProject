/*
*Auth: JackColor
*Date: 2019/1/5 下午12:03
*/
package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func getIcon(v string) int {

	sum := 0
	for _, v := range v {

		switch {
		case v == 'a', v == 'A':
			sum++
		case v == 'e', v == 'E':
			sum++
		case v == 'i', v == 'I':
			sum = sum + 2

		case v == 'o', v == 'O':
			sum = sum + 3
		case v == 'u', v == 'U':
			sum = sum + 5

		}

	}
	return sum
}

func distribute() int {

	left := 100
	for _, v := range users {
		sum := getIcon(v)
		left = left - sum
		val, ok := distribution[v]
		if !ok {
			distribution[v] = sum
		} else {

			distribution[v] = val + sum

		}

	}
	return left
}

func main() {

	left :=distribute()
	for k,v :=range distribution{
		fmt.Printf("%s get the icon %d\n",k,v)

	}
	fmt.Println(left)


}
