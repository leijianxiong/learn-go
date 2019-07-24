package main

import (
	"log"
	"net/http"
)

func main()  {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":10002", nil)
	if err != nil {
		log.Fatal("http listener err:", err)
	}
}

func index(w http.ResponseWriter, r *http.Request)  {
	_, err := w.Write([]byte("hello"))
	if err != nil {
		log.Fatal("rsp write err:", err)
	}
}
