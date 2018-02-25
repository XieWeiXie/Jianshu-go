package jianshu

type PassageDetail struct {
	time     string `json:"time"`
	title    string `json:"title"`
	abstract string `json:"abstract"`
	reader   int    `json:"reader"`
	comment  int    `json:"comment"`
	liked    int    `json:"liked"`
	payed    int    `json:"payed"`
	author   string `json:"author"`
}

type Topics struct {
	name string `json:"name"`
}

type SpecialTopics struct {
	name                string `json:"name"`
	url                 string `json:"url"`
	creator             string `json:"creator"`
	totalPassage        string `json:"total_passage"`
	totalFollowerNumber string `json:"total_follower_number"`
}

type LikeNotes struct {
	author    string `json:"author"`
	authorUrl string `json:"authorHomePage"`
	title     string `json:"title"`
	abstract  string `json:"abstract"`
}

// 分三类: 评论、喜欢、关注
type LatestActive struct {
	activeType   string `json:"type"`
	activeDetail string `json:"detail"`
}

// 最新评论的文章
type LatestCommented struct {
	title    string `json:"title"`
	abstract string `json:"abstract"`
	url      string `json:"url"`
}

// 首页的热门文集
type TopicCollection struct {
	name                  string `json:"name"`
	collectionDescription string `json:"description"`
	passageAndFollower    string `json:"detail"`
}
