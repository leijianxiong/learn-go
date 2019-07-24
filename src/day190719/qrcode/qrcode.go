package main

import "github.com/skip2/go-qrcode"

func main() {
	qrcode.WriteFile("http://www.flysnow.org/",qrcode.Medium,256,"./blog_qrcode.png")
}
