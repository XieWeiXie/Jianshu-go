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

type LatestComment struct {
	time         string `json:"time""`
	comment      string `json:"comment"`
	passageTitle string `json:"passage_title"`
}
