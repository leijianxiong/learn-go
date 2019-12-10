package main

import (
	"errors"
	"fmt"
)

/*
测试返回类型 不写返回值就表示没有返回值, 没有返回nil一说,
 */

func f1(name string) {
	fmt.Println("name:", name)
}

func main()  {
	f1("abc")
	errors.New("abc")
}
