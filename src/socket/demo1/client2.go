package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

/*
json发过来保持单行的话, 内容里面有空行如何处理!
{"a":"aa\nbb","b":"cc\ndd"}
 */

func main()  {
	conn, e := net.Dial("tcp", ":10001")
	log.Println("net dail e:", e)

	content := "{\"a\":\"aa\nbb\",\"b\":\"cc\ndd\"}"
	n, e := fmt.Fprintf(conn, content + "\nabc\n")
	log.Printf("conn write %d bytes", n, e)

	s, e := bufio.NewReader(conn).ReadString('\n')
	log.Println("conn read:", s, e)
}