package jianshu

import (
	"github.com/PuerkitoBio/goquery"
)

type Topic struct {
	url    string
	number int
	doc    *goquery.Document
}

func NewTopic(number int) *Topic {
	if number <= 0 {
		return nil
	}
	hotPageDoc, _ := goquery.NewDocument(rootUrl)
	hotCollection, _ := hotPageDoc.Find("a.more-hot-collection").Attr("href")
	url := MakeCompleteUrl(hotCollection)
	doc, _ := goquery.NewDocument(url)
	return &Topic{
		url:    url,
		number: number,
		doc:    doc,
	}
}

// 推荐
func (topic *Topic) GetTopicCollectionRecommend() []TopicCollection {
	url := topic.getURL("recommend", 0)
	//fmt.Println(topic.getURL("recommend", 0))
	return topic.getDocSelection(url)

}

// 热门
func (topic *Topic) GetTopicCollectionHot() []TopicCollection {
	url := topic.getURL("hot", 1)
	//fmt.Println(topic.getURL("hot", 1))
	return topic.getDocSelection(url)
}

// 城市
func (topic *Topic) GetTopicCollectionCity() []TopicCollection {
	url := topic.getURL("city", 2)
	//fmt.Println(topic.getURL("city", 2))
	return topic.getDocSelection(url)
}

// 校园
func (topic *Topic) GetTopicCollectionSchoolyard() []TopicCollection {
	url := topic.getURL("schoolyard", 3)
	//fmt.Println(topic.getURL("schoolyard", 3))
	return topic.getDocSelection(url)
}

func (topic *Topic) getURL(key string, index int) string {
	doc := topic.doc
	data, _ := doc.Find("ul.trigger-menu li").Eq(index).Find("a").Attr("href")
	return MakeCompleteUrl(data) + "&_pjax=%23list-container"
}

func (topic *Topic) getDocSelection(url string) []TopicCollection {
	var allTopicCollection []TopicCollection
	doc, _ := goquery.NewDocument(url)
	doc.Find("div.col-xs-8").Each(func(i int, selection *goquery.Selection) {
		if i < topic.number {
			name := selection.Find("div.collection-wrap a h4.name").Text()
			collectionDescription := StringCommon(selection.Find("div.collection-wrap a p.collection-description").Text())
			passageAndFollower := StringSpace(selection.Find("div.count").Text())
			//fmt.Println(name, collectionDescription, passageAndFollower)
			one := TopicCollection{
				name: name,
				collectionDescription: collectionDescription,
				passageAndFollower:    passageAndFollower,
			}
			allTopicCollection = append(allTopicCollection, one)
		}
	})
	return allTopicCollection

}
