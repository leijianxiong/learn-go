package main

import (
	"fmt"
	"sync"
)

/*
race.go => 初步判断为多个线程同时写同一块内存地址
race2.go => 声明了一个非并发安全字典 m, 并发线程里声明标量a,
=>不同线程里声明的a有自己的内存地址, 读取每个m["a"]的值也是访问相同的内存地址 所以暂时没有数据竞争
race3.go
	尝试不用锁来操作并发线程的变量 => 思路1: 传递结构体类型变量的地址 通过![√]
race4.go 测试同一个变量同时读写 会出现数据竞争的问题
 */


type Human struct {
	Name string
	age int
	gender Gender
}

type Gender int

func NewHuman(name string, age int) *Human {
	return &Human{
		Name: name,
		age: age,
		gender: GENDER_UNKNOWN,
	}
}

func (h *Human) SetName(name string) {
	h.Name = name
}

const GENDER_UNKNOWN Gender = 0
const GENDER_MALE Gender = 1
const GENDER_FAMALE Gender = 2

func main()  {
	wg := sync.WaitGroup{}
	wg.Add(5)

	human := NewHuman("a", 12)

	for i := 1; i <= 5; i++ {
		name := fmt.Sprintf("leijianxiong-%d", i)
		age := 25 + i
		go func(name string, age int, i int) {
			defer wg.Done()

			human.SetName(fmt.Sprintf("jx-%d", i))

			fmt.Printf("i=%d, human name=%s age=%d, human p:%p\n", i, human.Name, human.age, human)
		}(name, age, i)
	}

	wg.Wait()
	fmt.Println("main end")
}
