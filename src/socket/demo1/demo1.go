package main

import (
	"log"
	"net"
)

/*
	监听10001端口, 接收到消息后回复消息
 */

func main() {
	log.Println("begin listener..")
	listener, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Fatal("net.Listen err:", err)
	}
	log.Println("listener ok.")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("listener.Accept err:", err)
		}

		go handleAccept(conn)
	}
}

func handleAccept(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal("conn close err:", err)
		}
	}()
	log.Println("new conn start")

	//parse read
	//var req []byte
	//reader := bufio.NewReader(conn)
	//for {
	//	bytes, e := reader.ReadBytes('\n')
	//	log.Println("readbytes:", string(bytes), e)
	//	if e != nil {
	//		if e == io.EOF {
	//			req = append(req, bytes...)
	//			break
	//		}
	//		log.Fatal("reader read-bytes err:", e)
	//	}
	//	req = append(req, bytes...)
	//}
	//log.Println("req:", string(req))

	//parse read-v2
	var r []byte
	req := make([]byte, 5)
	for i := 1; i <= 15; i++ {
		n, err := conn.Read(req)
		if err != nil {
			log.Fatal("conn read err:", err)
		}
		log.Printf("conn read %d bytes content=%s", n, req[0:n], err)
		r = append(r, req[0:n]...)
		if i == 4 {
			//break
		}
	}
	log.Printf("[final] conn read %d bytes content=%s", len(r), string(r))

	//b := make([]byte, 5)
	//n, e := conn.Read(b)
	//if e != nil {
	//	log.Println("conn read err:", e)
	//}
	//log.Printf("read %d bytes content=\"%s\"", n, string(b))

	//req := make([]byte, 10)
	//i, e := conn.Read(req)
	//log.Printf("read %d bytes, content=\"%s\" err=%v", i, string(req), e)

	//rsp
	//rsp := "this is response \n content\nab\ncd2"
	//rsp := "this is response content ab cd2"
	rsp := "this is response content"
	n, err := conn.Write([]byte(rsp))
	if err != nil {
		log.Println("write err:", err)
	}
	log.Printf("write n(%d) %d bytes, content=\"%s\"", n, len(rsp), rsp)
	log.Println("new conn end")
}
