package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
)

/*
网络请求发送json串过去

避免内容有\n字符, 必要时base64 encode
 */

func main()  {
	conn, err := net.Dial("tcp", ":10001")
	if err != nil {
		log.Fatal("net dial err:", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal("conn close err:", err)
		}
		log.Println("conn closed.")
	}()

	data := map[string]string{
		"aa": "11",
		"bb": "2\n2",
	}
	//data := struct {
	//	AA string `json:"aa"`
	//	BB string `json:"bb"`
	//}{
	//	AA: "11",
	//	BB: "2\n2",
	//}
	//log.Println("data[bb]=", data.bb)
	//encode
	d, err := json.Marshal(data)
	if err != nil {
		log.Fatal("json encode err:", err)
	}
	log.Println("encoded:", string(d))

	//write
	d = append(d, '\n')
	n, err := conn.Write(d)
	if err != nil {
		log.Fatal("conn write err:", err)
	}
	log.Printf("conn write %v bytes, content=%v\n", n, string(d))

	//read
	rsp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("get rsp err:", err)
	}
	log.Println("rsp:", rsp)
}
