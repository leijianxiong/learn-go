package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

/*
把staruml导出的fragment json文件再处理, 输出格式化后的字符

检查参数个数
循环文件 处理每一个文件
	循环每一列 处理每一列
 */

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./format file1 [file2..]")
		os.Exit(1)
	}

	type Fragment struct {
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
	var fragment Fragment
	var lines string

	for _, filename := range os.Args[1:] {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("os openfile err:", err)
			return
		}

		decoder := json.NewDecoder(file)
		err = decoder.Decode(&fragment)
		if err != nil {
			fmt.Println("decode err:", err)
			return
		}

		fmt.Println(fragment.Name + "\t" + fragment.Documentation)
		defaultPattern, err := regexp.Compile(`d=\[(?P<default>.*)\]\s+|\n+`)
		if err != nil {
			fmt.Println("regexp compile err:", err)
		}
		for _, column := range fragment.Columns {
			var dataType, nullable, defaultValue, comment string
			dataType = column.Type
			switch column.Length.(type) {
			case float64:
				if column.Length != 0.0 {
					dataType += "("+fmt.Sprintf("%v", column.Length)+")"
				}
			case string:
				if column.Length != "" {
					dataType += "("+fmt.Sprintf("%v", column.Length)+")"
				}
			}
			if column.Nullable {
				nullable = "true"
			} else {
				nullable = "false"
			}
			match := defaultPattern.FindStringSubmatch(column.Documentation)
			if match != nil {
				defaultValue = match[1]
			} else {
				defaultValue = "null"
			}
			comment = defaultPattern.ReplaceAllString(column.Documentation, "")
			//$lines .= $column->name."\t$type\t nullable=$nullable\tdefault=$default\tcomment=\"$comment\"\n";
			lines += column.Name +"\t"+dataType+"\tnullable="+nullable+"\tdefault="+defaultValue+"\tcomment=\""+comment+"\"\n"
		}

		//call column -t
		cmd := "cat <<EOF | column -t\n"+lines+"EOF\n"
		output, err := exec.Command("bash", "-c", cmd).Output()
		if err != nil {
			fmt.Println("cmd output err:", err)
			return
		}
		fmt.Println(string(output))
		fmt.Println()
	}
}

