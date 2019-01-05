/*
*Auth: JackColor
*Date: 2019/1/1 上午11:17
*/
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func Sort() {

	var v [10]int = [10]int{1, 2, 3, 4, 5, 6, 6, 7, 8, 9}

	sort.Ints(v[:])
	fmt.Printf("sort int %d\n", v)

	var b [5]string = [5]string{"A", "B", "C", "D", "E"}

	sort.Strings(b[:])

	fmt.Printf("sort string %s\n", b)

}

var (
	length int
	class  string
)

const (
	NumStr  = "1234567890"
	CharStr = "zxcvbnmasdfghjklqwertyuiopZXCVBNMASDFGHJKLQWERTYUIO"
	SpeStr  = "/[]!@#$%^&*()_+~"
)

func ParserArgs() {

	flag.IntVar(&length,"-l",16,"长度")
	flag.StringVar(&class,"-t","num","类型")

	flag.Parse()

}

func generatePassword() string {

	password := make([]byte, length, length)

	var sourceStr string

	if class == "num" {
		sourceStr = NumStr

	} else if class == "char" {
		sourceStr = CharStr
	} else if class == "mix" {

		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpeStr)
	} else {
		sourceStr = NumStr
	}

	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		password[i] = sourceStr[index]
	}

	return string(password)

}

func main() {
	rand.Seed(time.Now().UnixNano())
	ParserArgs()
	fmt.Println(length,class)



	password := generatePassword()

	fmt.Printf("password %s\n", password)

}
