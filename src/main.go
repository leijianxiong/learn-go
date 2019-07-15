package main

import (
	"fmt"
	"time"
)

/*
计算fib的值
 */

func fib(n int, c chan<- int) {
	fmt.Println(time.Now())
	x, y := 1, 1
	for i := 1; i <= n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fib(10, c)
	for v := range c {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

