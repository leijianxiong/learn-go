package main

import (
	"fmt"
)

func main() {
	c := make(chan struct{}, 1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("i:", i)
		}
		c <- struct{}{}
	}()

	<- c
}
