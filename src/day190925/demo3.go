package main

import "fmt"

//select 先逐一求值候选分支表达式, 如果是阻塞的就表示是不成功的
// 该例子建立容量2的通道, 发送一个值, 然后关闭这个通道
//select位于for=3, 第一次会获取到值, 第二次在一个已经关闭了的通道中求值, 会得到这个通道元素类型的零值, 第二个返回值为false, 第三次等同第二次

func main()  {
	ch := make(chan int, 2)
	ch <- 1
	close(ch)

	for i := 0; i < 3; i++ {
		select {
		case v, ok := <- ch:
			fmt.Println("v=", v, "ok=", ok)
		default:
			fmt.Println("default case")
		}
	}
}