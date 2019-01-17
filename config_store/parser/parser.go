/*
*Auth: JackColor
*Date: 2019/1/14 下午9:26
*/
package parser

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
	"strings"
)

//序列化 结构体转换成 文件流
func Marshal(data interface{}) (res []byte, err error) {

	return

}

//文件转换成结构体
func UnMarshal(data []byte, result interface{}) (err error) {
	ResultType := reflect.TypeOf(result)

	if ResultType.Kind() != reflect.Ptr {

		err = errors.New("should pass a address")
		return
	}

	if ResultType.Elem().Kind() != reflect.Struct {

		err = errors.New("should pass a struct")
		return

	}

	lineArray := strings.Split(string(data), "\n")

	var lastFiledName string
	// 对应的结构体
	typeStruct := reflect.TypeOf(result).Elem()

	for index, value := range lineArray {
		_ = index
		line := strings.TrimSpace(value)
		if len(line) == 0 {
			continue
		}
		if line[0] == ';' || line[0] == '#' {
			continue
		}
		//fmt.Println(index, value)

		if line[0] == '[' && len(line) <= 2 {
			err = errors.Errorf("syntax error, invalid section:%s", line)
		}

		if line[0] == '[' && line[len(line)-1] != ']' {
			err = errors.Errorf("syntax error, invalid section:%s", line)

		}

		if line[0] == '[' && line[len(line)-1] == ']' {
			//解析 每一个 sectionName
			lastFiledName, err = parserSections(line, typeStruct)
			//fmt.Println(lastFiledName)
			continue

			if err != nil {
				err = fmt.Errorf("%v,lineNo:%d", err, index+1)
				return
			}

		}
		// 把 sectionName 放进去 去拿到这个Name对应的结构体  然后拿到这个结构体的Name 赋值
		err = parserFields(lastFiledName, line, result)

		if err != nil {
			err = errors.Errorf("%v,LineNo:%d", err, index+1)
			return
		}

	}

	return
}
func parserFields(lastFiledName, line string, result interface{}) (err error) {

	// 解析每一行数据

	index := strings.Index(line, "=")
	if index == -1 {

		err = errors.Errorf("the line without = line:%s", line)
		return
	}

	key := strings.TrimSpace(line[0:index])
	value := strings.TrimSpace(line[index+1:])
	//_ = value
	if len(key) == 0 {

		err = errors.Errorf("the line without key map line:%s", line)
		return

	}
	//得到key value
	//fmt.Println(key, value)

	//找到 结构体对应的所有的值
	resultValue := reflect.ValueOf(result)

	// 找到lastFiledName 对应值 对用的 Value
	sectionValue := resultValue.Elem().FieldByName(lastFiledName)
	//找到 找到lastFiledName 的type 必须是一个结构体
	sectionType := sectionValue.Type()

	if sectionType.Kind() != reflect.Struct {

		err = fmt.Errorf("the section type must sturct %s", lastFiledName)
		return

	}
	var KeyName string
	for i := 0; i < sectionType.NumField(); i++ {

		filed := sectionType.Field(i)
		tagVal := filed.Tag.Get("ini")
		if tagVal == key {
			KeyName = filed.Name
			break
		}

	}
	if len(KeyName) == 0 {

		return
	}
	//
	keyVal := sectionValue.FieldByName(KeyName)
	if keyVal == reflect.ValueOf(nil) {
		return
	}
	switch keyVal.Kind() {

	case reflect.String:
		keyVal.SetString(value)

	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		PaserVal, errRes := strconv.ParseInt(value, 10, 64)
		if errRes != nil {
			err = errRes
			err = fmt.Errorf("can't transfer the string to int")
			return
		}
		keyVal.SetInt(PaserVal)

		// 浮点数,无符号整型

	}

	return

}
func parserSections(line string, typeStruct reflect.Type) (lastFiledName string, err error) {

	sectionName := line[1 : len(line)-1]
	sectionName = strings.TrimSpace(sectionName)
	if len(sectionName) == 0 {
		err = errors.Errorf("syntax error, invalid section:%s", line)
	}

	for i := 0; i < typeStruct.NumField(); i++ {

		filed := typeStruct.Field(i)

		tagName := filed.Tag.Get("ini") //获得结构体的ini  对应的字段

		if sectionName == tagName {
			//fmt.Println(filed.Name)
			//把 filed name 赋值给 最外的变量`
			lastFiledName = filed.Name
			return
		}
	}

	return
}
