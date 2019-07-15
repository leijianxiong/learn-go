package main

import (
	"bytes"
	"fmt"
)

/*

 */

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("abc\ndef\n123\n45"))
	fmt.Println("buffer中内容:", buffer.String())
	line, err := buffer.ReadString('\n')
	fmt.Println("ReadString result:", line, err)
	fmt.Println("buffer中内容:", buffer.String())
}
