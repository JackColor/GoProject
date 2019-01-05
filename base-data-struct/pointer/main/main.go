/*
*Auth: JackColor
*Date: 2019/1/1 下午12:24
*/
package main

import "fmt"

func main()  {


   var a int = 32


   var c *int

   if c==nil{
   	fmt.Printf("c is nil ==>%p\n",c)
   }



   b :=&a

   fmt.Printf("b is a pointer from a the addr is %p\n",b)

   //修改b内存地址 指向的值

   *b ++

   fmt.Printf("b 修改后的值 %d\n",*b)
   fmt.Printf("a 的值 %d",a)









}




