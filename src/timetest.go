package main

import (
	"fmt"
	"time"
)

func main() {
	var ts0, ts1 int64
	ts0 = time.Now().Unix()
	time.Sleep(time.Second * 3)
	ts1 = time.Now().Unix()
	fmt.Printf("ts0=%d, ts1=%d ts1-ts0=%d\n", ts0, ts1, ts1-ts0)
}
