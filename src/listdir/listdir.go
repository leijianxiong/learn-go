package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/**
给定一个目录 列出目录里所有文件
 */

var entry string
var maxLevel int

func init() {
	flag.StringVar(&entry, "entry", "", "entry usage")
	flag.IntVar(&maxLevel, "max-level", 1, "max-level usage")
}

func main()  {
	flag.Parse()
	fmt.Println("entry:", entry)
	if entry == "" {
		panic("entry require. usage: -entry=dir")
	}

	if err := listdir(entry, 1); err != nil {
		panic(err)
	}
}

func listdir(entry string, currentLevel int) error {
	if currentLevel > maxLevel {
		//log.Println("level=max")
		return nil
	}

	files, err := ioutil.ReadDir(entry)
	if err != nil {
		log.Fatal(err)
		return err
	}
	for _, file := range files {
		if file.Name() == "." || file.Name() == ".." {
			continue
		}
		filename := strings.TrimRight(entry, "/") + "/" + file.Name()
		if file.IsDir() {
			fmt.Println(filename + "/")
			_ = listdir(filename, currentLevel + 1)
		} else {
			fmt.Println(filename)
		}

	}
	return nil
}

