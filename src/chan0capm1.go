package main

import (
	"fmt"
	"time"
)

/**
发送方向通道发送值
接收方接收
 */

func main() {
	sendingInterval := time.Second
	receptionInterval := time.Second * 2
	//intChan := make(chan int, 0)
	intChan := make(chan int, 5)
	go func() {
		var t0, t1 int64
		for i := 1; i <= 5; i++ {
			t0 = time.Now().Unix()
			intChan <- i
			if i == 1 {
				fmt.Println("Sent:", i)
			} else {
				time.Sleep(sendingInterval)
				t1 = time.Now().Unix()
				fmt.Printf("Sent: %d [interval: %d s]\n", i, t1-t0)
			}
		}
		close(intChan)
	}()

	var t0, t1 int64
Loop:
	for {
		t0 = time.Now().Unix()
		select {
		case v, ok := <-intChan:
			if !ok {
				break Loop
			}
			if v == 1 {
				fmt.Println("Received:", v)
			} else {
				time.Sleep(receptionInterval)
				t1 = time.Now().Unix()
				fmt.Printf("Received: %d [interval: %d s]\n", v, t1-t0)
			}
		}
	}
	fmt.Println("End.")


}
