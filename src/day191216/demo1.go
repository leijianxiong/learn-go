package main

import (
	"fmt"
	"sync"
)

func main()  {
	//m := sync.Map{}

	wg := sync.WaitGroup{}
	wg.Add(20)

	fmt.Println("print start")

	for i := 0; i <= 9; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf(" %d ", i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Printf(" %d ", i)
		}(i)
	}

	wg.Wait()
	fmt.Println("\nprint end")
}
