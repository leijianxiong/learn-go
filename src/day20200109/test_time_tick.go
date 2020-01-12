package main

import (
	"fmt"
	"time"
)

func main()  {
	ticker := time.NewTicker(1 * time.Second)

	time.AfterFunc(5 * time.Second, func() {
		fmt.Println("ticker stop")
		ticker.Stop()
	})

	loop:
	for {
		select {
		case t, ok := <-ticker.C:
			if !ok {
				fmt.Println("ticker chan closed, break")
				break loop
			}
			fmt.Println("t is:", t, "do something")
		default:
			fmt.Println("ticker stoped")
			break loop
		}
	}

	fmt.Println("for end")


}
