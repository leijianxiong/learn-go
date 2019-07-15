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
	sendingInterval := time.Second * 4
	receptionInterval := time.Second * 2
	//intChan := make(chan int, 0)
	intChan := make(chan int, 5)
	go func() {
		var t0, t1 int64
		for i := 1; i <= 5; i++ {
			intChan <- i
			t1 = time.Now().Unix()
			if t0 == 0 {
				fmt.Println("Sent:", i)
			} else {
				fmt.Printf("Sent: %d [interval: %d s]\n", i, t1-t0)
			}
			t0 = time.Now().Unix()
			time.Sleep(sendingInterval)
		}
		close(intChan)
	}()

	var t0, t1 int64
Loop:
	for {
		select {
		case v, ok := <-intChan:
			if !ok {
				break Loop
			}
			t1 = time.Now().Unix()
			if t0 == 0 {
				fmt.Println("Received: ", v)
			} else {
				fmt.Printf("Received: %d [interval: %d s]\n", v, t1-t0)
			}
		}
		t0 = time.Now().Unix()
		time.Sleep(receptionInterval)
	}
	fmt.Println("End.")


}
