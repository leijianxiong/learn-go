package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main()  {
	listener, e := net.Listen("tcp", ":10001")
	log.Println("net listen err:", e)

	for {
		conn, e := listener.Accept()
		log.Println("listener accept err:", e)

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn)  {
	defer conn.Close()
	s, e := bufio.NewReader(conn).ReadString('\n')
	log.Println("conn read ", s, e, strings.ContainsAny(s, "\n"))

	n, e := fmt.Fprintf(conn, "ok\n")
	log.Printf("conn write %d bytes, err:%s", n, e)
}
