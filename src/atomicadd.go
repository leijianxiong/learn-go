package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var ui32 uint32 = 100
	fmt.Println("ui32:", ui32)
	atomic.AddUint32(&ui32, ^uint32(0))
	fmt.Println("ui32:", ui32)
	fmt.Printf("^uint32(0):%d uint32(1):%d\n", ^uint32(0), uint32(1))
}
