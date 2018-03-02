package jianshu

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

func (a *Article) GetAuthor() {}

func (a *Article) GetDescription() {}

func (a *Article) GetTitle() {}

func (a *Article) GetContent() {}

func (a *Article) GetInfo() {}
