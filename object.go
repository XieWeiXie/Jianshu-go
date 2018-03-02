package jianshu

type PassageDetail struct {
	time     string
	title    string
	abstract string
	reader   int
	comment  int
	liked    int
	payed    int
	author   string
}

type Topics struct {
	name string
}

type SpecialTopics struct {
	name                string
	url                 string
	creator             string
	totalPassage        string
	totalFollowerNumber string
}

type LikeNotes struct {
	author    string
	authorUrl string
	title     string
	abstract  string
}

// 分三类: 评论、喜欢、关注
type LatestActive struct {
	activeType   string
	activeDetail string
}

// 最新评论的文章
type LatestCommented struct {
	title    string
	abstract string
	url      string
}

// 首页的热门文集
type TopicCollection struct {
	name                  string
	collectionDescription string
	passageAndFollower    string
}

// 首页的推荐作者
type Author struct {
	name        string
	url         string
	gender      string
	description string
	newDetail   string
}

// 首页热门文章
type HotPassage struct {
	author         string
	time           string
	collectionName string
	reader         string
	comment        string
	liked          string
	payed          string
	title          string
	url            string
	abstract       string
}

// 出版书籍

type Book struct {
	author   string
	title    string
	abstract string
	price    string
	url      string
}
