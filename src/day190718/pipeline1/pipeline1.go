package main

import (
	"fmt"
	"os/exec"
)

/*
测试管道操作

参考 /Users/jianxiong/Documents/staruml/format.php

给各个文件格式化输出
 */

func main()  {
	demo1()

}

/*
简单测试 使用管道
 */
type item struct {
	name string
	age int
}

func demo1()  {

	data := []item{
		{
			name: "a",
			age: 10,
		},
		{
			name: "bb",
			age: 11,
		},
		{
			name: "ccc",
			age: 13,
		},
		{
			name: "dddd",
			age: 14,
		},
	}
	str := fmt.Sprintf("Index Name Age\n")

	for k, v := range data {
		//fmt.Printf("k=%v, v=%v\n", k, v)
		str += fmt.Sprintf("%d %s %d\n", k, v.name, v.age)
	}

	fmt.Println("str")
	fmt.Println(str)

	//format
	cmd := "cat <<EOF | column -t\n"+str+"EOF\n"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println("command output err:", err)
	}
	fmt.Println("output-format:")
	fmt.Println(string(out))
}
