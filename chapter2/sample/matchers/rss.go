package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/goinaction/code/chapter2/sample/search"
)

type (
	// item结构体
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image结构体
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}

	// channel结构体
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument结构体
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)

// rssMatcher本身是一个结构体，但是由于它实现了Search方法，所以可以直接使用
type rssMatcher struct{}

// 接口和结构体是一个东西，前提是该结构体实现了接口所需要的函数

// 唯一被调用的地方就是在main函数中隐式引入的
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// 为结构定义方法，不是在结构体内部，而是创建好了结构体之后，单独应该一个函数，该函数名前传入结构体别名和结构体名称即可。
// 为结构体定义了一个Search函数
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	// 获取已经结构化的数据
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// 遍历数据中的每个Item，正则查询title
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// 找到结果后传递给通道
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// 遍历数据中的每个Item，正则查询Description
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// 找到结果后传递给通道
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}

// 为结构体定义了一个retrieve函数
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	// 接收一个结构体，输出一个结构体和一个错误信息
	if feed.URI == "" {
		return nil, errors.New("No rss feed uri provided")
	}

	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// 结束后关闭
	defer resp.Body.Close()

	// 验证响应码
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	// 定义匹配的结构进行由于接收返回值
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
