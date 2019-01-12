/*
*Auth: JackColor
*Date: 2019/1/12 下午2:07
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ListDir(path string, deep int) (err error) {

	file, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	if deep == 1 {

		fmt.Printf("!----%s\n", filepath.Base(path))
	}
	sep := string(os.PathSeparator)
	for _, val := range file {

		if val.IsDir() {
			fmt.Printf("|")

			for i := 0; i < deep; i++ {
				fmt.Printf("    |")
			}
			fmt.Printf("----%s\n", val.Name())

			ListDir(path+sep+val.Name(), deep+1)

			continue
		}

		fmt.Printf("|")
		for i := 0; i < deep; i++ {
			fmt.Printf("    |")
		}

		fmt.Printf("-----%s\n", val.Name())

	}

	return

}

func main() {

	ListDir("./", 1)

	//args 模块 cli 使用

}
