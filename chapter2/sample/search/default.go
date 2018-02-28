package search

// 空结构体
type defaultMatcher struct{}

// 引入search包的时候会自动被引入
func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// 同样定义了一个Search方法
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
