package search

import (
	"log"
)

// Result是一个结构体，也是通道元素的基本组成
type Result struct {
	Field   string
	Content string
}

// Matcher是一个接口，里面包含一个Search方法是实际查询的方式
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// 每个Goroutine实际上执行的内容，所有结构体前面需要加*
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 执行查询
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 将查询得到的结果传送回通道中
	for _, result := range searchResults {
		results <- result
	}
}

func Display(results chan *Result) {
	// 通道关闭后，for循环才会停止，否则会一直运行
	// 实时打印通道中的数据
	for result := range results {
		log.Printf("IIIII%s:\n%s\n\n", result.Field, result.Content)
	}
}
