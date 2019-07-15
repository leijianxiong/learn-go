package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
)


func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息 fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, "")) }
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

//登录
func login(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("/Users/jianxiong/work/go/learn-go/src/goweb/demo3.2/login.html")
		if err != nil {
			fmt.Println("template parse files err:", err)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			fmt.Println("execute get error", err)
		}
	} else {
		r.ParseForm()

		fmt.Println("handle form post", r.PostForm, r.FormValue("username"), r.PostFormValue("username"),
			r.URL.Query().Get("username"))
		fmt.Println("form:", r.Form)
		fmt.Println("username:", r.Form["username"], "username type", reflect.TypeOf(r.Form["username"]))
		fmt.Println("password:", r.Form["password"])
	}
}

func upload(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles("learn-go/src/goweb/demo3.2/upload.html")
		if err != nil {
			fmt.Println("template.ParseFiles err:", err)
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			fmt.Println("t.execute err:", err)
			return
		}
	} else {
		r.ParseMultipartForm(4*1024*1024)
		file, handler, err := r.FormFile("uploadfile")
		defer file.Close()
		if err != nil {
			fmt.Println("r.form-file err:", err)
			return
		}
		fmt.Fprintf(w, "%v", handler.Header)
		//strings.explode
		filenameSplits := strings.Split(handler.Filename, "/")
		f, err := os.OpenFile("/tmp/"+filenameSplits[len(filenameSplits)-1], os.O_WRONLY|os.O_CREATE, 0666)
		//f, err := os.OpenFile("/tmp/5ab0d13c85162.jpg", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("[error] http.ListenAndServe: ", err)
	}
}
