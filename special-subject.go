package jianshu

import (
	"github.com/PuerkitoBio/goquery"
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

func (s *SpecialSubject) GetSpecialSubjectTitleAdministrator() {}

func (s *SpecialSubject) GetSpecialSubjectNewComment() {}

func (s *SpecialSubject) GetSpecialSubjectNewAdd() {}

func (s *SpecialSubject) GetSpecialSubjectHot() {}

func (s *SpecialSubject) getResponse() {

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
