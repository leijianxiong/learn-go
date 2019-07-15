package main

import "fmt"

func main() {
	var v1 *m1 = genI()
	var v2 *m1 = genI()
	fmt.Printf("v1=%p, v2=%p\n", v1, v2)
	fmt.Println("执行修改操作")
	v1.a++
	fmt.Printf("v1.a=%d, v2.a=%d\n", v1.a, v2.a)

}

type m1 struct {
	a int
	b int
}

func genI() *m1  {
	return &m1{
		a: 0,
		b: 0,
	}
}
