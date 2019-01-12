/*
*Auth: JackColor
*Date: 2019/1/12 下午6:57
*/
package log

import "fmt"

type FileLog struct {
	path string
}

func NewFileLogHandler(path string) logInterface {

	return &FileLog{
		path: path,
	}

}

func (f *FileLog) LogWarn() {

	fmt.Printf("print warn log in to the file %s\n", f.path)
}

func (f *FileLog) LogDebug() {
	fmt.Printf("print debug log in to the file %s\n", f.path)
}
