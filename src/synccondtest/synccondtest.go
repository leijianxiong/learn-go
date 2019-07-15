package main

import "sync"

/*
测试条件变量
 */

func main() {
	var mailbox uint8
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(lock.RLocker())

}
