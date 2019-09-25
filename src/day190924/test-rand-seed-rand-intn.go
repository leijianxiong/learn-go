package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("time now nano=", t.UnixNano())
	fmt.Println("time now=", t.Unix())

	//每次程序运行都要保证随机数不一样
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 30; i++ {
		// [0,n)
		v := rand.Intn(30)
		fmt.Println("i=" , i, " v=", v)
	}
}
