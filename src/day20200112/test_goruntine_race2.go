package main

import (
	"fmt"
	"sync"
)

/*
race.go => 初步判断为多个线程同时写同一块内存地址
race2.go => 声明了一个非并发安全字典 m, 并发线程里声明标量a,
=>不同线程里声明的a有自己的内存地址, 读取每个m["a"]的值也是访问相同的内存地址 所以暂时没有数据竞争
 */
func main()  {
	wg := sync.WaitGroup{}
	wg.Add(5)

	m := map[string]int{
		"a": 1,
	}

	for i := 1; i <= 5; i++ {
		go func(i int) {
			defer wg.Done()

			a := 0
			fmt.Printf("i=%d a=%d ap=%p\n", i, a, &a)

			a = i + m["a"]
			fmt.Printf("i=%d a=%d ap=%p added\n", i, a, &a)
		}(i)
	}

	wg.Wait()
	fmt.Println("main end")
}
