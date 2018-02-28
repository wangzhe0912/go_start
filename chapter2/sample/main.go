package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers" // _ 表示仅仅引用，调用了init函数，但是没有显示调用
	"github.com/goinaction/code/chapter2/sample/search"
)

// 在main函数之前执行
func init() {
	// 设置输出格式
	log.SetOutput(os.Stdout)
}

// 程序入口
func main() {
	// 执行查询
	search.Run("president")
}
