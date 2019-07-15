package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)
	closed := make(chan struct{}, 1)
	go func() {
		ticker := time.NewTicker(time.Millisecond * 500)
		defer ticker.Stop()
	LOOP:
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			case <-closed:
				break LOOP
			}
		}
		done <- true
	}()

	time.Sleep(time.Millisecond * 1600)
	closed <- struct{}{}
	<-done
	fmt.Println("Ticker stopped")
}