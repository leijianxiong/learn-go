package main

import (
	"errors"
	"fmt"
)

/*
进入一个函数
开头处加入defer
	defer部分处理
下面引发panic
 */

func main() {
	fmt.Println("enter main func")
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("defer recover e(%s)", e)
		}
	}()

	panic(errors.New("main func panic call1"))
	fmt.Println("exit main func")
}
