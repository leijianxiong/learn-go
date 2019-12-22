package main

import "fmt"

func main() {
	//值类型
	var file = struct {
		Name string
	}{
		Name:"abc",
	}
	fmt.Printf("file: %s\n", file.Name)

	//指针类型
	var file2 = new(struct{
		Name string
	})
	file2.Name = "a2"
	fmt.Printf("file2: %s\n", file2.Name)
	fmt.Printf("file2-2: %s\n", (*file2).Name)

}
