package main

import (
	"html/template"
	"log"
	"os"
)

/*
测试template parsefile
 */
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	gopath, err := os.Getwd()
	checkErr(err)

	log.Println("dir:", gopath)

	t, err := template.ParseFiles(gopath + "/src/day190719/http/login.gtpl")
	checkErr(err)
	err = t.Execute(os.Stdout, nil)
	checkErr(err)
}
