/*
*Auth: JackColor
*Date: 2019/1/16 下午11:45
*/
package do_test

import (
	"study/_test/demo"
	"testing"
)

func TestDemo(t *testing.T)  {

	a:=1
	b:=2
	c:= demo.Add(a,b)
	if c!=3{

		t.Fatal("failed to test the Add")
	}
	t.Logf("success test add")




}