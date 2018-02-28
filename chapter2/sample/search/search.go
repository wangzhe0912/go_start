package search

import (
	"log"
	"sync"
)

// 创建一个Map，其key为字符串形式，value为一个Matcher接口
var matchers = make(map[string]Matcher)

// 主要的查询函数，接收一个字符串参数
func Run(searchTerm string) {
	// 直接调用feed.go下的RetrieveFeeds函数来读取文件
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲的通道用于将获取到的匹配结果打印出来
	results := make(chan *Result)

	// 创建一个计数器
	var waitGroup sync.WaitGroup

	// 计数器的值设置为切片的长度
	waitGroup.Add(len(feeds))

	// 遍历切片中的每个元素，针对每个元素启动一个Goroutine处理匹配任务。
	for _, feed := range feeds {
		// 找到对应的匹配匹配器
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动Goroutine处理匹配任务，匿名函数
		go func(matcher Matcher, feed *Feed) {
			// 调用Match函数，分别传入了matcher接口，其中包含Search方法，feed数据，searchTerm内容以及接收数据的通道results
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done() //每个任务执行完成后，计数器减少1
		}(matcher, feed) //直接传入参数运行
	}

	// 同时启动一个Goroutine用于等待任务全部完成后关闭通道。
	go func() {
		// 等待任务全部执行完成
		waitGroup.Wait()
		// 用于终止程序，如果通道不关闭，程序I不会结束。
		close(results)
	}()

	// 实时打印通道中的数据
	Display(results)
}

// 用于注册matcher，传入的参数分别是matcher的名称和接口（通用函数）
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		// 不允许重名注册
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
