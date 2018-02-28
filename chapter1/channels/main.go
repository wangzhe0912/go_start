package main

import (
	"encoding/json"
	"os"
	"fmt"
)

const file = `c:\Go\src\github.com\goinaction\code\chapter2\sample\data\data1.json`

type MyDataStruct struct {
	// Struct中的变量如果没有大写，那么在Struct外部全部都是不可见的
	// 赋值变量 类型 解析的格式
	Name string `json:"name"`
	Age  int `json:"age"`
	Sex string `json:"sex"`
}

func load_json_file(filename string) ([]*MyDataStruct, error) {
	file_content, _ := os.Open(file)
	defer file_content.Close()
	var mydata []*MyDataStruct
	error := json.NewDecoder(file_content).Decode(&mydata)
	return mydata, error
}

func main() {
	// 执行查询
	mydata, _ := load_json_file(`c:\Go\src\github.com\goinaction\code\chapter2\sample\data\data1.json`)
	fmt.Println(len(mydata))
	var value = mydata[0].Sex
	fmt.Println(value)
	fmt.Println(mydata)  //直接打印一个切片，则看到将会是一个地址。
	fmt.Println(mydata[0])
	fmt.Println(mydata[0].Name)
}
