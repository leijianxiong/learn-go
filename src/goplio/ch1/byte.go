package main

import (
	"fmt"
	"reflect"
)

func main() {
	var b []byte = []byte("你好啊")
	fmt.Println(reflect.TypeOf(b[0]))
}
