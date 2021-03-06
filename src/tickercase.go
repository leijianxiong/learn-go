package main

import (
	"fmt"
	"time"
)

/*
	发送方使用断续器向一个通道发送值
	接收方获取这个值并累计, 满足条件后关闭断续器
 */

func main() {
	var ticker *time.Ticker = time.NewTicker(time.Millisecond * 500)
	intChan := make(chan int, 1)
	sendEndChan := make(chan struct{}, 1)
	doneChan := make(chan struct{}, 1)
	go func() {
	Loop:
		for _ = range ticker.C {
			select {
			case intChan <- 1:
			case intChan <- 2:
			case intChan <- 3:
			case intChan <- 4:
			case <-sendEndChan:
				break Loop
			}
		}
		doneChan <- struct {}{}
		fmt.Println("End. [Sender]")
		defer ticker.Stop()
	}()

	sum := 0
	for i := range intChan {
		fmt.Println("Received:", i)
		sum += i
		if sum > 10 {
			fmt.Printf("Got %d, break\n", sum)
			sendEndChan <- struct{}{}
			break
		}
	}
	<-doneChan

	fmt.Println("End. [receiver]")
}
