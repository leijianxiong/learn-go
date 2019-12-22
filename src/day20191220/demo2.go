package main

/*
测试并发获取db数据! 看值会不会覆盖
 */

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)
import _ "github.com/go-sql-driver/mysql"

func main()  {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.33.22:3306)/www_abohui_com")
	if err != nil {
		log.Fatal(err)
	}

	type User struct {
		uid int
		account string
	}

	user := new(User)
	user.uid = 1
	fmt.Printf("user %+v %d\n", *user, user.uid)

	usersChan := make(chan *User, 6)

	wg := sync.WaitGroup{}
	wg.Add(3)

	perPage := 2
	for i := 1; i <= 3; i++ {
		go func(page int) {
			defer wg.Done()
			offset := (page - 1) * perPage
			sql := fmt.Sprintf("select uid,account from eb_user order by uid asc limit %d,%d", offset, perPage)
			fmt.Printf("page=%d sql=%s\n", page, sql)
			rows, err := db.Query(sql)
			if err != nil {
				log.Fatalf("page=%d err=%s", page, err)
			}

			i := int32(0)
			for rows.Next() {
				user := new(User)
				err := rows.Scan(&user.uid, &user.account)
				if err != nil {
					log.Fatalf("page=%d err=%s", page, err)
				}

				fmt.Printf("page=%d i=%d user(%p) uid=%d uid-p=%p account=%s account-p=%p\n",
					page, i, user, user.uid, &user.uid, user.account, &user.account)
				usersChan <- user
				i = atomic.AddInt32(&i, 1)
			}
		}(i)
	}

	wg.Wait()

	close(usersChan)

	fmt.Println("read from userschan:")
	for user := range usersChan {
		fmt.Printf("user=%+v uid=%d account=%s\n", *user, user.uid, user.account)
		//users := append(users, user)
	}
	fmt.Println("end.")
}

