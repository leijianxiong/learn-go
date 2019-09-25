package main

import (
	"fmt"
	"unsafe"
)

//测试数据类型 nil
/*
nil
bool int float string byte(uint8) 这里说的是字节切片 []byte
array slice map channel function
struct interface pointer
 */

func main() {
	//fmt.Println(time.Now().String())
	var b bool
	fmt.Printf("%v\n", b) //false

	var i int
	fmt.Printf("%v\n", i) //0

	var f1 float64
	var f2 float64 = 0.0
	fmt.Printf("%v, f1==f2?:%v\n", f1, f1==f2) //0, true

	var s1 string
	fmt.Printf("s1=\"%v\"\n", s1) //s1=""

	var b1 byte
	fmt.Printf("b1=%v\n", b1) //b1=0

	var b2 []byte
	fmt.Printf("b2=%v, %v\n", b2, b2==nil) //b2=[], true

	var arr1 [3]int
	fmt.Printf("arr1=%v\n", arr1) //数组会分配个数, 每个元素是该元素类型的零值

	var m1 map[string]int
	fmt.Printf("map:%v, m1==nil?:%v\n", m1, m1 == nil)

	var chan1 chan int
	fmt.Printf("chan:%v\n", chan1) //chan:<nil>
	var chan2 = make(chan int, 2)
	chan2 <- 10
	chan2 <- 20
	fmt.Printf("chan2:%v, %v, %v, %v\n", chan2, <-chan2, <-chan2, nil)
	//对满了元素的通道进行发送元素时会报错
	//对没有元素的通道进行接收元素时会报错
	//所以对通道的接收操作使用select语句

	var func1 func(a int, b string) int
	fmt.Printf("func1=%v\n", func1) //func1=<nil>

	/*
	pointers -> nil
	slices -> nil
	maps -> nil
	channels -> nil
	functions -> nil
	interfaces -> nil
	 */

	var p1 unsafe.Pointer
	fmt.Printf("p1=%v\n", p1) //p1=<nil>

	var if1 interface{}
	fmt.Printf("if1=%v", if1)
}
