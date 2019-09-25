package main

/**
go操作redis
 */

import (
	"fmt"
	"github.com/piaohao/godis"
)

func main() {
	redis := godis.NewRedis(&godis.Option{
		Host: "192.168.40.16",
		//Host: "192.168.10.128",
		Port: 6379,
		Db:   0,
	})
	defer redis.Close()

	//string
	redis.Set("godis", "1")
	//fmt.Printf("s=%v, err:%v\n", s, err)
	s1, _ := redis.Get("godis")
	println(s1)

	//列表
	ms := []string{"aa", "bb", "cc"}
	redis.LPush("list-key1", ms...)

	for {
		s, _ := redis.RPop("list-key1")
		println("s=", s)
		if s == "" {
			println("s=\"\", break.")
			break
		}
	}

	//hash hset hget (hmset, hmget) hdel
	//hset
	redis.HSet("hash-key1", "k1", "v1")
	m, _ := redis.HGetAll("hash-key1")
	fmt.Println("hash-key1=", m)

	//hget
	k1value, _ := redis.HGet("hash-key1", "k1")
	println("k1:", k1value)

	m2 := map[string]string{
		"k2": "v2",
		"k3": "v3",
	}

	redis.HMSet("hash-key1", m2)
	s3, _ := redis.HMGet("hash-key1", []string{"k3", "k2"}...)
	fmt.Println("s3:", s3)
}