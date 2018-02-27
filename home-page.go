package jianshu

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type HomePage struct {
	number int
	page   int
}

// number <= 20 page <= 10
func NewHomePage(number, page int) *HomePage {
	if page <= 1 {
		page = 1
	}
	if number > 20 {
		return nil
	}
	return &HomePage{
		number: number,
		page:   page,
	}
}

func (h *HomePage) doc() *goquery.Document {
	doc, _ := goquery.NewDocument(rootUrl)
	return doc

}

func (h *HomePage) GetHomePagePassages() []HotPassage {
	var allHotPassageCollections []HotPassage
	doc := h.doc()
	doc.Find("ul.note-list li div.content").Each(func(i int, selection *goquery.Selection) {
		if i < h.number {
			var collectionName, reader, comment, payed string
			author := selection.Find("div.author a.nickname").Text()
			tempTime, _ := selection.Find("span.time").Attr("href")
			time := StringSpace(tempTime)
			title := StringSpace(selection.Find("a.title").Text())
			tempUrl, _ := selection.Find("a.title").Attr("href")
			url := MakeCompleteUrl(tempUrl)
			abstract := StringSpace(StringCommon(selection.Find("p.abstract").Text()))
			flagCollection, _ := selection.Find("div.meta a").Eq(0).Attr("class")
			if strings.Contains(flagCollection, "collection-tag") {
				collectionName = StringSpace(selection.Find("div.meta a.collection-tag").Text())
			}
			sizeA := selection.Find("div.meta").Find("a").Size()
			if sizeA == 3 {
				reader = StringSpace(selection.Find("div.meta a").Eq(1).Text())
				comment = StringSpace(selection.Find("div.meta a").Eq(2).Text())
			}
			sizeSpan := selection.Find("div.meta span").Size()
			liked := StringSpace(selection.Find("div.meta span").Eq(0).Text())
			if sizeSpan == 2 {
				payed = StringSpace(selection.Find("div.meta span").Eq(1).Text())
			}
			one := HotPassage{
				author:         author,
				time:           time,
				title:          title,
				url:            url,
				abstract:       abstract,
				collectionName: collectionName,
				reader:         reader,
				comment:        comment,
				liked:          liked,
				payed:          payed,
			}
			//fmt.Println(author, time, title, url, abstract, collectionName, reader, comment, liked, payed)
			allHotPassageCollections = append(allHotPassageCollections, one)
		}
	})
	return allHotPassageCollections
}

func (h *HomePage) GetNewList() {}

func (h *HomePage) GetHotSeven() {}

func (h *HomePage) GetHotMonth() {}

func (h *HomePage) GetJianShuSchool() {}
