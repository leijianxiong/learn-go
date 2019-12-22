package main

import (
	"fmt"
	"sync"
)

/*
测试并发代码里创建引用变量会不会是同一个值?

 */

type Int1 int

func NewInt1() *Int1 {
	return new(Int1)
}

func main()  {

	// 1 测试打印变量指针

	//i := 1
	//i2 := i
	//fmt.Printf("i 地址为 %p \n", &i)
	//fmt.Printf("i2 地址为 %p \n", &i2)
	//
	//ia := &i
	//fmt.Printf("i=%d ia指针地址为%p i地址为%p\n", *ia, &ia, ia)
	//
	////2测试引用类型
	//s := []int{1, 2}
	//s2 := s
	//fmt.Printf("s地址为%p 第一个元素%d地址为%p\n", &s, s[0], &s[0])
	//fmt.Printf("s2地址为%p\n", &s2)
	//fmt.Printf("s2 第一个元素%d地址%p\n", s2[0], &s2[0])

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i <= 9; i++ {
		go func(i int) {
			defer wg.Done()
			a1 := NewInt1()
			fmt.Printf("new int1 得到变量a1, 值为%d, 地址为%p\n", *a1, a1)
		}(i)
	}

	wg.Wait()
	fmt.Println("end.")
}
