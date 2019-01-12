/*
*Auth: JackColor
*Date: 2019/1/6 下午2:39
*/
package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) > 1 {

		for index, val := range os.Args[1:] {

			fmt.Printf("os.args[%d]=%s\n", index, val)

		}

	}

}
