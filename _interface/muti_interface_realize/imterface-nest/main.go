/*
*Auth: JackColor
*Date: 2019/1/12 下午6:20
*/
package main

import "fmt"

type Read interface {
	read()
}

type Write interface {
	write()
}

// 接口继承
type File interface {
	Read
	Write
}

type Files struct {
	path string
	size int
}

func (f *Files) read() {

	fmt.Printf("the file path is %s and the size is %d\n", f.path, f.size)

}

func (f *Files) write() {

	fmt.Println("the file can write")

}

func main() {
	var f *Files = &Files{path: "./test", size: 1000}

	var r Read
	r = f
	r.read()

	var a File

	a = f

	a.read()
	a.write()

}
