package jianshu

import (
	"encoding/json"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

type SpecialSubject struct {
	URL string
}

func NewSpecialSubject(url string) *SpecialSubject {
	if url == "" {
		return nil
	}
	return &SpecialSubject{
		URL: url,
	}

}

func (s *SpecialSubject) doc() *goquery.Document {
	doc, _ := goquery.NewDocument(s.URL)
	return doc
}

func (s *SpecialSubject) GetSpecialSubjectURL() string {
	return s.URL
}

func (s *SpecialSubject) GetSpecialSubjectTitle() string {
	doc := s.doc()
	return StringSpace(doc.Find("div.title").Text())
}

func (s *SpecialSubject) GetSpecialSubjectPassageNumber() int {
	return s.getNumber(0)
}

func (s *SpecialSubject) GetSpecialSubjectFollowerNumber() int {
	return s.getNumber(1)
}

func (s *SpecialSubject) GetSpecialSubjectNotice() string {
	doc := s.doc()
	return StringCommon(doc.Find("div.description.js-description").Text())
}

func (s *SpecialSubject) GetSpecialSubjectTitleAdministrator() interface{} {
	doc := s.doc()
	collectionString := doc.Find(`script[data-name="collection"]`).Text()
	collections := &collection{}
	err := json.Unmarshal([]byte(collectionString), collections)
	if err != nil {
		return nil
	}
	url := GetSpecialSubjectURL(collections.Id)
	newDoc, _ := goquery.NewDocument(url)
	resultSlug := MakeCompleteUrl("u/" + gjson.Get(newDoc.Text(), "editors.#.slug").Array()[0].String())
	resultNickName := gjson.Get(newDoc.Text(), "editors.#.nickname").Array()[0].String()
	newCreator := creator{
		URL:      resultSlug,
		NickName: resultNickName,
	}
	return newCreator
}

func (s *SpecialSubject) GetSpecialSubjectNewComment() []PassageDetail {
	return s.getResponse("commented")
}

func (s *SpecialSubject) GetSpecialSubjectNewAdd() []PassageDetail {
	return s.getResponse("added")
}

func (s *SpecialSubject) GetSpecialSubjectHot() []PassageDetail {
	return s.getResponse("top")
}

func (s *SpecialSubject) getNumber(index int) int {
	doc := s.doc()
	numberOne, numberTwo := RegexpFindSting(StringSpace(doc.Find("div.main-top div.info").Text()))
	if index == 0 {
		return numberOne
	}
	if index == 1 {
		return numberTwo
	}
	return 0
}

func (s *SpecialSubject) getURL(key string) string {
	if key == "commented" {
		return s.URL + "?order_by=commented_at&_pjax=%23list-container"
	}
	if key == "top" {
		return s.URL + "?order_by=top&_pjax=%23list-container"
	}
	if key == "added" {
		return s.URL + "?order_by=added_at&_pjax=%23list-container"
	}
	return "None"
}
func (s *SpecialSubject) getResponse(key string) []PassageDetail {
	var allPassageDetail []PassageDetail
	url := s.getURL(key)
	doc, _ := goquery.NewDocument(url)
	doc.Find("ul.note-list li").Each(func(i int, selector *goquery.Selection) {
		var onePassage PassageDetail
		time, exist := selector.Find("div.content").Find("div.author").Find("div.info span").Attr("data-shared-at")
		if !exist {
			time = ""
		}
		title := selector.Find("div.content").Find("a.title").Text()
		abstract := StringSpace(selector.Find("div.content").Find("p").Text())
		author := StringSpace(selector.Find("div.content").Find("div.author div.info").Text())
		reader := StringToInt(StringSpace(selector.Find("div.content div.meta").Find("a").Eq(0).Text()))
		comment := StringToInt(StringSpace(selector.Find("div.content div.meta").Find("a").Eq(1).Text()))
		liked := StringToInt(StringSpace(selector.Find("div.content div.meta").Find("span").Eq(0).Text()))
		payed := StringToInt(StringSpace(selector.Find("div.content div.meta").Find("span").Eq(1).Text()))
		onePassage.title = title
		onePassage.abstract = abstract
		onePassage.author = author
		onePassage.reader = reader
		onePassage.comment = comment
		onePassage.liked = liked
		onePassage.payed = payed
		onePassage.time = time
		allPassageDetail = append(allPassageDetail, onePassage)
	})
	return allPassageDetail
}

type collection struct {
	Id               int            `json:"id"`
	Slug             string         `json:"slug"`
	Owner            map[string]int `json:"owner"`
	SubscribersCount int            `json:"subscribers_count"`
}

type creator struct {
	URL      string `json:"url"`
	NickName string `json:"nickname"`
}
