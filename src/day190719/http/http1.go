package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

/*
测试表单提交
 */

func checkErr(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}

func main()  {
	http.HandleFunc("/", Index)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	checkErr(err)
}

func Index(w http.ResponseWriter, r *http.Request )  {
	_, err := w.Write([]byte("hello"))
	checkErr(err)
}

func login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	checkErr(err)
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		//entry, err := os.Getwd()
		//checkErr(err)
		//t, _ := template.ParseFiles(entry + "/src/day190719/http/login.gtpl")
		t, _ := template.ParseFiles("/Users/jianxiong/work/go/learn-go/src/day190719/http/login.gtpl")
		err := t.Execute(w, nil)
		checkErr(err)

	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username from url query:", r.URL.Query().Get("username"))
		fmt.Println("username from Post-Form:", r.PostForm["username"])
		fmt.Println("password from Post-Form:", r.PostForm["password"])
	}
}
