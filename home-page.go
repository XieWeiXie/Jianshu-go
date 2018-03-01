package jianshu

import (
	"strings"

	"fmt"

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

func (h *HomePage) GetNewList() {
	h.getResponse(0)

}

func (h *HomePage) GetHotSeven() {
	h.getResponse(1)
}

func (h *HomePage) GetHotMonth() {
	h.getResponse(2)

}

func (h *HomePage) GetJianShuSchool() {
	h.getResponse(4)
}

func (h *HomePage) getURL(index int) string {
	var url string
	if index > 5 {
		return "None"
	}
	doc := h.doc()
	size := doc.Find("div.board").Find("a").Size()
	if size >= 4 {
		tempUrl, _ := doc.Find("div.board").Find("a").Eq(index).Attr("href")
		url = MakeCompleteUrl(tempUrl)
	}
	return url

}

func (h *HomePage) getResponse(index int) []PassageDetail {
	/*
		    time     string
			title    string
			abstract string
			reader   int
			comment  int
			liked    int
			payed    int
			author   string
	*/
	var allPassageDetail []PassageDetail
	url := h.getURL(index)
	doc, _ := goquery.NewDocument(url)
	doc.Find("div#list-container ul.note-list li").Each(func(i int, selection *goquery.Selection) {
		var reader, comment, liked, payed string
		author := StringSpace(selection.Find("div.content div.author div.info a.nickname").Text())
		time := StringSpace(selection.Find("div.content div.author div.info span.time").Text())
		title := StringSpace(selection.Find("div.content div.author a.title").Text())
		abstract := StringCommon(selection.Find("div.content p.abstract").Text())
		if selection.Find("div.content div.meta a").Size() == 3 {
			reader = StringCommon(selection.Find("div.content div.meta a").Eq(1).Text())
			comment = StringCommon(selection.Find("div.content div.meta a").Eq(2).Text())
		} else if selection.Find("div.content div.meta a").Size() == 2 {
			reader = StringCommon(selection.Find("div.content div.meta a").Eq(0).Text())
			comment = StringCommon(selection.Find("div.content div.meta a").Eq(1).Text())
		} else {
			reader = StringCommon(selection.Find("div.content div.meta a").Eq(0).Text())
			comment = "0"
		}
		if selection.Find("div.content div.meta span").Size() == 2 {
			liked = StringCommon(selection.Find("div.content div.meta span").Eq(0).Text())
			payed = StringCommon(selection.Find("div.content div.meta span").Eq(1).Text())
		} else {
			liked = StringCommon(selection.Find("div.content div.meta span").Eq(0).Text())
			payed = "0"
		}

		fmt.Println(author, time, title, abstract, reader, comment, liked, payed)
	})
	return allPassageDetail
}
