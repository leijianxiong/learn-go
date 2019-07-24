//这个文件是这个包总的例子介绍
//每个函数下的例子介绍使用单独的Example`FuncName`[_param1]
package godocexample

import (
	"fmt"
)

type a struct {
	Aa string
	Bb string
}

func Example() {
	sum:=Add(1,2)
	fmt.Println("1+2=",sum)
	//output Example

	//Output:
	//1+2=3
}

//add doc
func ExampleAdd() {
	//add 1+2
	i := Add(1, 2)
	fmt.Println("1+2=", i)

	//Output:
	//1+2=3
}

// The fmt package's Errorf function lets us use the package's formatting
// features to create descriptive error messages.
func ExampleAdd_errofparam() {
	//test example with p1 p2
	Add(1, 3)

	//Output:
	//1+3=4
}