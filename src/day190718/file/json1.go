package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
处理json

 */


func main()  {
	//demo2()
	//demo3()
	//demo4()
	demo5()
}

func demo5()  {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}

func demo4()  {
	var jsonBlob = []byte(`[
	{"name": "Platypus", "order": "Monotremata"},
	{"name": "Quoll",    "order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

func demo2()  {
	//var str = `{"aaData":[["018002633552","\u8d44\u6599\u5f55\u5165","routeAdd","<a class=\"btn default btn-xs purple\" href=\"\/admin\/workflow\/editunit?wid=018002633552\" ><i class=\"fa fa-edit\"><\/i> \u67e5\u770b\u7f16\u8f91<\/a>"],["014305616747","\u8d44\u6599\u5ba1\u6838","viewCustomer","<a class=\"btn default btn-xs purple\" href=\"\/admin\/workflow\/editunit?wid=014305616747\" ><i class=\"fa fa-edit\"><\/i> \u67e5\u770b\u7f16\u8f91<\/a>"],["464201340364","\u4e1a\u52a1\u7533\u8bf7","bzapply","<a class=\"btn default btn-xs purple\" href=\"\/admin\/workflow\/editunit?wid=464201340364\" ><i class=\"fa fa-edit\"><\/i> \u67e5\u770b\u7f16\u8f91<\/a>"],["198801423759","\u4e1a\u52a1\u5f00\u901a\u5ba1\u6838","bzaudit","<a class=\"btn default btn-xs purple\" href=\"\/admin\/workflow\/editunit?wid=198801423759\" ><i class=\"fa fa-edit\"><\/i> \u67e5\u770b\u7f16\u8f91<\/a>"]],"totalNum":4}`
	var postsStr = `[["018002633552","\u8d44\u6599\u5f55\u5165","routeAdd","<a class=\"btn default btn-xs purple\" href=\"\/admin\/workflow\/editunit?wid=018002633552\" ><i class=\"fa fa-edit\"><\/i> \u67e5\u770b\u7f16\u8f91<\/a>"],["014305616747","\u8d44\u6599\u5ba1\u6838","viewCustomer","<a class=\"btn default btn-xs purple\" href=\"\/admin\/workflow\/editunit?wid=014305616747\" ><i class=\"fa fa-edit\"><\/i> \u67e5\u770b\u7f16\u8f91<\/a>"],["464201340364","\u4e1a\u52a1\u7533\u8bf7","bzapply","<a class=\"btn default btn-xs purple\" href=\"\/admin\/workflow\/editunit?wid=464201340364\" ><i class=\"fa fa-edit\"><\/i> \u67e5\u770b\u7f16\u8f91<\/a>"],["198801423759","\u4e1a\u52a1\u5f00\u901a\u5ba1\u6838","bzaudit","<a class=\"btn default btn-xs purple\" href=\"\/admin\/workflow\/editunit?wid=198801423759\" ><i class=\"fa fa-edit\"><\/i> \u67e5\u770b\u7f16\u8f91<\/a>"]]`
	var f [][]string
	buffer := bytes.NewBufferString(postsStr)

	err := json.Unmarshal([]byte(postsStr), &f)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}

	fmt.Println("测试json decode 多维数组:")
	for k, v := range f {
		fmt.Printf("k=%v\n", k)
		for k1, v1 := range v {
			fmt.Printf("\tk1=%v, v1=%v\n", k1, v1)
		}
	}


	fmt.Println("测试json decode多维数据 采用json.Decoder")
	decoder := json.NewDecoder(buffer)
	err = decoder.Decode(&f)
	if err != nil {
		fmt.Println("decode err:", err)
	}
	fmt.Printf("subf=%v\n", f)
}

type Mfj struct {
	Type_ string `json:"_type"`
	Id_ string `json:"_id"`
	Parent_ map[string]string `json:"_parent"`
	Name string `json:"name"`
	Documentation string `json:"documentation"`
	Columns []struct{
		Type_ string `json:"_type"`
		Id_ string `json:"_id"`
		Parent_ map[string]string `json:"_parent"`
		Name string `json:"name"`
		Documentation string `json:"documentation"`
		Type string `json:"type"`
		Length interface{} `json:"length"`
		PrimaryKey bool `json:"primaryKey"`
		Nullable bool `json:"nullable"`
	} `json:"columns"`
}

/*
int
int32
int64
float32
float64
to
string
 */
func ParseString(length interface{}) string {
	return fmt.Sprintf("%v", length)
}

func demo3()  {
	fmt.Println("\n使用文件方式读取复杂json内容")
	if len(os.Args) < 2 {
		log.Fatal("err: filepath required.")
		return
	}
	filename := os.Args[1]
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("openfile err:", err)
		return
	}

	decoder := json.NewDecoder(file)
	var mfj Mfj
	err = decoder.Decode(&mfj)
	if err != nil {
		fmt.Println("decode err:", err)
		return
	}
	fmt.Println("mfj:", mfj)

	//处理字符不同
	cc := len(mfj.Columns)
	cs := mfj.Columns[cc-4:cc]
	for k, v := range cs {
		s := ParseString(v.Length)
		fmt.Printf("k=%v v=%v, s=%v\n", k, "", s)
	}
}
