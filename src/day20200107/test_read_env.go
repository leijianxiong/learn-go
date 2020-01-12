package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

/*
测试.env
 */

func main()  {
	//f, err := os.OpenFile("./.env", os.O_RDONLY, os.ModePerm)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//s, _ := ioutil.ReadFile("./.env")
	s, err := ioutil.ReadFile("./.env")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("ioutil.ReadFile: ", string(s))

	f, err := os.Open("./.env")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	fmt.Println("bufio.newReader: ")

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("line:", line)
	}

}


