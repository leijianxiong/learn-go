package main

import "fmt"

/*
v1 先实现这个
v2 改成异步通道也是可以的吧
 */
func fib2(c chan<- int, quit <-chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("fib2 default selected")
		}
	}
}

func main() {
	c := make(chan int, 10)
	quit := make(chan int, 1)
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib2(c, quit)
}
