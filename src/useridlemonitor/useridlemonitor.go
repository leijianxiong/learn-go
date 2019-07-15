package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/mattn/go-oci8"
	"log"
	"os"
	"time"
)

/*
	入口 命令行执行
		命令行参数加默认值
	oracle查询数据
	获取到条件更新数据
 */


func main() {
	var day int
	flag.IntVar(&day, "day", 1, "user idle max time")
	fmt.Println("args day:", day)

	fmt.Println("os.args=", os.Args, len(os.Args))
	//oracle 查询数据
	if len(os.Args) != 2 {
		fmt.Printf("ERROR: Please provide a DSN string in ONE argument:\n\n")
		fmt.Println("Shell-Conversion into DSN string:")
		fmt.Println("  sqlplus sys/password@tnsentry as sysdba   =>   sys/password@tnsentry?as=sysdba")
		fmt.Println("  sqlplus / as sysdba                       =>   sys/.@?as=sysdba")
		fmt.Println("instead of the tnsentry, you can also use the hostname of the IP.")
		os.Exit(1)
	}
	os.Setenv("NLS_LANG", "")

	//var dsn = "oci:dbname=192.168.40.15:1521/tjpay;charset=utf"
	tns := "crm_user/mMsp1234@192.168.40.15:1521/tjpay"
	db, err := sql.Open("oci8", tns)
	if err != nil {
		fmt.Println("sql.open err:", err)
		return
	}
	defer db.Close()
	fmt.Println()
	var user string
	//var users []string
	idlemaxtime := 30*24*3600
	sqlstring := fmt.Sprintf("select user_id from crm_user.t_admusers where status='1' AND to_number(LASTLOGINTIME) < %d",
		time.Now().Unix()-int64(idlemaxtime))
	log.Print("sql:", sqlstring)

	err = db.QueryRow(sqlstring).Scan(&user)
	//rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("db.query-row scan err:", err)
		return
	}

	//if err := rows.Scan(&users); err != nil {
	//	log.Println(err)
	//}

	fmt.Printf("Successful 'as sysdba' connection. Current user is: %v\n", user)

}