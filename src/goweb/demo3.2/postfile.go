package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

/*
模拟发图片请求
/Users/jianxiong/Pictures/5ab0d13c85162.jpg

 */

/*
创建buffer
模拟一个文件buffer
 */
func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	fmt.Println("content-type:", contentType)
	bodyWriter.Close()
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func main()  {
	filename := "/Users/jianxiong/Pictures/5ab0d13c85162.jpg"
	targeturl := "http://localhost:9090/upload"
	err := postFile(filename, targeturl)
	fmt.Println("err:", err)
}


