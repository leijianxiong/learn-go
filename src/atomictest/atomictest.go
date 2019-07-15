package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var n int32 = 10
	swapped := atomic.CompareAndSwapInt32(&n, 10, 5)
	fmt.Printf("swapped %v n=%d", swapped, n)

	var v atomic.Value
}
