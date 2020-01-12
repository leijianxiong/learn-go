package main

import (
	"log"
	"os"
	"time"
)

/*
测试计划任务
 */
func main()  {
	//fmt.Println("t1")
	f, err := os.OpenFile("/tmp/test-crontab", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	_, err = f.WriteString("t update2 " + time.Now().String() + "\n")
	if err != nil {
		log.Fatal(err)
	}
}
