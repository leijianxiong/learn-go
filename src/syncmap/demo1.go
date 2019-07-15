package main

import (
	"fmt"
	"sync"
)

/*
测试sync.map
 */

func main() {
	m := sync.Map{}
	m.Store("a", 1)
	m.Store("b", 2)

	a, _ := m.Load("a")
	fmt.Println("a=", a)
	b, _ := m.Load("b")
	fmt.Println("b=", b)
}
