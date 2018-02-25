package jianshu

import (
	"github.com/PuerkitoBio/goquery"
)

type RecommendAuthor struct {
	number int
}

func NewRecommendAuthor(number int) *RecommendAuthor {
	if number <= 0 {
		panic("number should larger than zero.")
	}
	if number > 24 {
		panic("number should less than twenty-four.")
	}
	return &RecommendAuthor{
		number: number,
	}
}

func (author *RecommendAuthor) GetListRecommendAuthor() []Author {
	return author.getResponse()
}

func (author *RecommendAuthor) getResponse() []Author {
	var allRecommendAuthor []Author
	url := HomePageAuthorUrl()
	doc, _ := goquery.NewDocument(url)
	doc.Find("div.col-xs-8").Each(func(i int, selection *goquery.Selection) {
		if i < author.number {
			var gender string
			name := StringCommon(selection.Find("div.wrap a h4.name").Text())
			tempUrl, _ := selection.Find("div.wrap a").First().Attr("href")
			url := MakeCompleteUrl(tempUrl)
			if selection.Find("div.wrap a h4 i").Size() == 0 {
				gender = "None"
			} else {
				tempGender, _ := selection.Find("div.wrap a h4 i").Attr("class")
				gender = SplitString(tempGender)
			}
			description := StringCommon(selection.Find("div.wrap p.description").Text())
			newDetail := StringCommon(selection.Find("div.recent-update").Text())
			//fmt.Println(name, url, gender, description, newDetail)
			one := Author{
				name:        name,
				url:         url,
				description: description,
				newDetail:   newDetail,
				gender:      gender,
			}
			allRecommendAuthor = append(allRecommendAuthor, one)
		}
	})
	return allRecommendAuthor
}
