package main

import (
	"log"
	"net"
	"os"
)

//client1

func main()  {
	log.Println("os args:", os.Args)

	log.Println("begin dail..")
	//conn, err := net.Dial("tcp", ":10001")
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:10001")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("net.Dail err:", err)
	}
	defer func() {
		e := conn.Close()
		if e != nil {
			log.Fatal("conn Close err:", e)
		}
		log.Println("client conn closed.")
	}()
	log.Println("dail ok.")

	n, err := conn.Write([]byte(os.Args[1]))
	if err != nil {
		log.Fatal("conn write err:", err)
	}
	log.Printf("conn write %d bytes, write content is \"%s\"", n, os.Args[1])

	//parse read


	b := make([]byte, 5)
	i, err := conn.Read(b)
	if err != nil {
		log.Println("conn read err:", err)
	}
	log.Printf("n=%d content=%s", i, string(b))

}
