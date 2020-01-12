package main

import "fmt"

/*
测试别名类型
 */

func main()  {
	is := make([]int, 0)
	fmt.Println("is", len(is))
	is = append(is, 1)
	fmt.Println("is", len(is), is)
}

type Int1 int

type M2 map[string]Int1

func f1(a M2, b Int1)  {

}
