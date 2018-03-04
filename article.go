package jianshu

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Article struct {
	url string
}

func NewArticle(url string) *Article {
	if url == "" {
		return nil
	}
	return &Article{
		url: url,
	}
}

func (a *Article) doc() *goquery.Document {
	doc, _ := goquery.NewDocument(a.url)
	return doc
}

// 获取作者
func (a *Article) GetAuthor() string {
	doc := a.doc()
	return StringCommon(doc.Find("div.article div.author div.info span.name").Text())
}

// 获取作者简介
func (a *Article) GetDescription() string {
	doc := a.doc()
	tempURL, _ := doc.Find("div.article div.author div.info span a").Attr("href")
	url := MakeCompleteUrl(tempURL)
	descriptionDoc, _ := goquery.NewDocument(url)
	tempDescription := StringSpace(descriptionDoc.Find("div.description div.js-intro").Text())
	return tempDescription
}

// 获取文章标题
func (a *Article) GetTitle() string {
	doc := a.doc()
	return doc.Find("div.article h1").Text()

}

// 获取文章内容
func (a *Article) GetContent() string {
	doc := a.doc()
	return strings.TrimSpace(doc.Find("div.article div.show-content").Text())

}
