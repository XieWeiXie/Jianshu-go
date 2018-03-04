package jianshu

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type TrendSearch struct {
}

func NewTrendSearch() *TrendSearch {
	return &TrendSearch{}
}

func (t *TrendSearch) doc() goquery.Document {
	doc, _ := goquery.NewDocument(GetTrendSearch())
	return *doc
}

// javascript
func (t *TrendSearch) GetTrendSearchList() {
	doc := t.doc()
	fmt.Println(doc.Html())

}
