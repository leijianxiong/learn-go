package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(11)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			go func(i int) {
				defer wg.Done()
				fmt.Println("i:", i)
			}(i)
		}
	}()

	wg.Wait()
}
