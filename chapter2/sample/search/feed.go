package search

import (
	"encoding/json"
	"os"
)

// Go语言在Windows下读取文件的方式
const dataFile = `c:\Go\src\github.com\goinaction\code\chapter2\sample\data\data.json`

// Feed contains information we need to process a feed.
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// 读取文件函数
func RetrieveFeeds() ([]*Feed, error) {
	// Open the file.
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 设置当文件读取完成后关闭文件
	defer file.Close()

	// 创建一个结构体变量，该结构体和JSON的格式一致，用于接收JSON的数据
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// 返回结果
	return feeds, err
}
