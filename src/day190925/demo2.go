package main

import "fmt"

//select 先逐一求值候选分支表达式, 如果是阻塞的就表示是不成功的
// 该例子建立容量2的通道, 发送一个值
//select位于for=3, 第一次会获取到值, 第二次时第一个case会阻塞, 会进入default默认分支, 第三次跟第二次一样

func main()  {
	ch := make(chan int, 2)
	ch <- 1

	for i := 0; i < 3; i++ {
		select {
		case v, ok := <- ch:
			fmt.Println("v=", v, "ok=", ok)
		default:
			fmt.Println("default case")
		}
	}
}
