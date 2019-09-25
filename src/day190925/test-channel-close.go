package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	//测试关闭一个channel后, select会进入这个分支?

	stopCh := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			select {
			case v, ok := <-stopCh:
				fmt.Println("receive v=", v, "ok=", ok)
				if !ok {
					fmt.Println("return")
					return
				}
			default:
				fmt.Println("default branch")
			}
		}
	}()

	//stopCh <- 1
	go func() {
		defer wg.Done()
		fmt.Println("sleep 1 second")
		time.Sleep(time.Microsecond * 1)
		//close(stopCh)
		//fmt.Println("closed stop-ch")
	}()

	wg.Wait()
}
