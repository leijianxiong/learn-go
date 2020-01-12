package main

import (
	"fmt"
	"time"
)

/*
通道第二个参数 只有在通道关闭时才会不ok
 */

func main()  {

	ch := make(chan int, 2)

	time.AfterFunc(1 * time.Second, func() {
		fmt.Println("ch close")
		close(ch)
	})

	fmt.Println("for start")
	loop:
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				fmt.Println("select ch closed:", i)
				break loop
			}
		default:
			fmt.Println("select default ")
		}

		fmt.Println("seelp 200 Millisecond ")
		time.Sleep(200 * time.Millisecond)
	}

}
