package jianshu

import (
	"strings"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type Publication struct {
}

func NewPublication() *Publication {
	return &Publication{}
}

func (p *Publication) doc() *goquery.Document {
	doc, _ := goquery.NewDocument(p.getURL())
	return doc
}

func (p *Publication) GetPublicizedBook() {
	p.getResponse("paper")
}

func (p *Publication) GetNovelBooks() {
	p.getResponse("novel")
}

func (p *Publication) GetITAndJobMarket() {
	p.getResponse("it")
}

func (p *Publication) GetCultureAndHistory() {
	p.getResponse("culture")
}

func (p *Publication) GetMonthlyMagazine() {
	p.getResponse("collection_monthly")
}

func (p *Publication) getURL() string {
	var url string
	doc, _ := goquery.NewDocument(rootUrl)
	doc.Find("div.board a").Each(func(i int, selection *goquery.Selection) {
		flag, _ := selection.Attr("href")
		if strings.Contains(flag, "/publications") {
			url = MakeCompleteUrl(flag)
		}
	})
	return url

}

func (p *Publication) getResponse(keyWord string) []Book {
	/*
		author   string
		title    string
		abstract string
		price    string
		url      string
	*/
	var allBook []Book
	doc := p.doc()
	// body > div.container.publications > div.col-xs-19.col-xs-offset-1.main
	doc.Find(fmt.Sprintf("div.col-xs-19.col-xs-offset-1.main div#%s a", keyWord)).Each(func(i int, selection *goquery.Selection) {
		var abstract string
		var one Book
		if selection.HasClass("book") {
			url, _ := selection.Attr("href")
			title := selection.Find("div.name").Text()
			author := selection.Find("div.author").Text()
			price := selection.Find("div.price").Text()
			one.url = url
			one.title = title
			one.author = author
			one.price = price
			if selection.Find("div").HasClass("intros") {
				abstract = selection.Find("div.intros").Text()
			} else {
				abstract = "None"
			}
			one.abstract = abstract
			fmt.Println(url, abstract, title, author, price)
			allBook = append(allBook, one)
		}

	})
	return allBook

}
