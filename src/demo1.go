package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "everyone", "usage string")

func init() {
	//flag.StringVar(&name, "name", "everyone", "usage string")
}

func main() {
	flag.Parse()
	fmt.Printf("hello %s!\n", *name)
}
