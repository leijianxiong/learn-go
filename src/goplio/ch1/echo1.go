package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//fmt.Println("args:", os.Args)

	var t0, t1 time.Time

	t0 = time.Now()
	for i, e := range os.Args {
		fmt.Println("arg:",i, e)
	}
	t1 = time.Now()
	fmt.Printf("t1-t0=%d s", t1.UnixNano()-t0.UnixNano())
	fmt.Println(t1)
	fmt.Println(t0)
}
