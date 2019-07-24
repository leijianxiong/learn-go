package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
读取文件内容
 */

func main()  {
	filename := os.Args[1]
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("openfile err:", err)
	}
	defer file.Close()

	//var content bytes.Buffer
	//_, err = io.Copy(&content, file)
	//if err != nil {
	//	log.Fatal("copy err:", err)
	//}
	//log.Println("content:", content.String())

	//fileBuf := bufio.NewReader(file)
	//for {
	//	line, isPrefix, err := fileBuf.ReadLine()
	//	fmt.Printf("line=%v, isprefix=%v, err=%v\n", string(line), isPrefix, err)
	//	//fmt.Println("line", line)
	//	//fmt.Println("isPrefix", isPrefix)
	//	//fmt.Println("err", err)
	//	if err != nil {
	//		fmt.Println("readline err:", err)
	//	}
	//	if !isPrefix {
	//		break;
	//	}
	//}

	demo3(file)
}

func demo3(file *os.File)  {
	fmt.Println("\n使用bufio reader.readstring处理单行文字")
	reader := bufio.NewReader(file)
	for {
		s, err := reader.ReadString('\n')
		s = strings.TrimRight(s, "\n")
		fmt.Println("s", s, "err", err)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read-string err:", err)
		}
	}
}
