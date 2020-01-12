package main

import "fmt"

func main()  {
	a, b := f1()
	fmt.Printf("a:%v, b:%v\n", a, b)
}

/*
返回值个数必须相等
 */
func f1() (a int, b error)  {
	return 1
}
