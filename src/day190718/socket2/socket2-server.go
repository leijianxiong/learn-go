package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
)

/*
网络请求解析json串
 */

func main()  {
	listener, err := net.Listen("tcp", ":10001")
	if err != nil {
		log.Fatal("net listen err:", err)
	}
	log.Println("Server start...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("listener accept err:", err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn)  {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal("conn close err:", err)
		}
		log.Println("new conn closed.")
	}()
	log.Println("\nnew conn start")
	req, err := bufio.NewReader(conn).ReadBytes('\n')
	if err == io.EOF {
		log.Println("readstring EOF");
		return
	}
	if err != nil {
		log.Fatal("reader readstring err:", err)
	}
	req = req[:len(req)-1]
	log.Printf("read %v bytes, content=%v\n", len(req), string(req))

	//返回这些信息
	//var reqFormat struct{
	//	AA string `json:"aa"`
	//	BB string `json:"bb"`
	//}
	//默认首字母大小写转换, 不一致的需要指定 json tag
	var reqFormat struct{
		Aa string
		Bb string
	}
	err = json.Unmarshal(req, &reqFormat)
	if err != nil {
		log.Fatal("json decode err:", err)
	}
	log.Println("req decoded:", reqFormat)

	_, err = conn.Write([]byte("ok\n"))
	if err != nil {
		log.Fatal("conn write err:", err)
	}
	log.Println("new conn end")

}
