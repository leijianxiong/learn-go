package main

import (
	"fmt"
	"sync"
)

/*
初步判断为多个线程同时写同一块内存地址
 */
func main()  {
	wg := sync.WaitGroup{}
	wg.Add(5)

	m := map[string]int{
		"a":1,
		"b":2,
	}

	for i := 1; i <= 5; i++ {
		go func(i int) {
			defer wg.Done()

			m["a"] = i
			a := m["a"]

			fmt.Printf("go-runtine:%d, a=%d, c=%d\n", i, a, m["b"])
		}(i)
	}

	wg.Wait()
	fmt.Println("main end")
}

func f2(a int) int {
	a += 1
	return a
}