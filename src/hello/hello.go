package main

import (
	"fmt"
)

func main() {
	slice := []string{"a ", "b "}
	testFunc(&slice)

	fmt.Println(slice)
}

func testFunc(slice *[]string) {
	(*slice)[0] = "111"
	fmt.Println(*slice)
}
