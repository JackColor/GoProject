/*
*Auth: JackColor
*Date: 2019/1/16 下午11:49
*/
package do_bench

import (
	"study/_test/demo"
	"testing"
)

func BenchmarkAdd(b *testing.B)  {

	// go test -bench .
	for i:=0;i<b.N;i++{

		a:=1
		b:=1
		demo.Add(a,b)



	}


}




