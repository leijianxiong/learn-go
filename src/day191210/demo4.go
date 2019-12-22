package main

import (
	"fmt"
)

/*
测试异常
 */

func f1() {
	defer func() {
		//step1 这里可以捕获f1的panic
		if p := recover(); p != nil {
			fmt.Printf("panic[defer 1]: %s\n", p)
			//这里重新抛出这个异常 或者不预期的结果出现异常 step2
			panic(p)
		}
	}()
	//step3 最后一个defer捕获前面defer出现的异常
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("panic[defer 2]: %s\n", p)
		}
	}()
	panic("f1 panic")
}

func main()  {
	f1()
}
