/*
*Auth: JackColor
*Date: 2019/1/5 下午7:54
*/
package main

import "fmt"

func testScan() {

	var a int
	var b string

	//fmt.Scanf("%d%s",&a,&b) //空格分割
	//
	//fmt.Printf("a=%d,b=%s",a,b)

	fmt.Scanf("%d\n", &a)

	fmt.Scanf("%s\n", &b)

	fmt.Printf("a=%d,b=%s", a, b)

}

func TestScan()  {
	var a int
	var b string
	fmt.Scan(&a,&b)
	fmt.Printf("a=%d,b=%s", a, b)

}

func testScanln()  {

	var a int
	var b string
	fmt.Scanln(&a,&b)
	fmt.Printf("a=%d,b=%s", a, b)


}

func testSscan()  {
	str := "11 aa"
	var a int
	var b string
	fmt.Sscan(str,&a,&b)

	fmt.Printf("a=%d,b=%s", a, b)

}

func testSscanf()  {
	str := "11 aa"
	var a int
	var b string

	fmt.Sscanf(str,"%d%s",&a,&b)

	fmt.Printf("a=%d,b=%s", a, b)

}

func test()  {
	str := "11aa"
	var a int
	var b string
	fmt.Sscanln(str,&a,&b)

	fmt.Printf("a=%d,b=%s", a, b)

}





func main() {

	//testScan()
	//TestScan()
	//testScanln()
	//testSscan()
	test()
}
