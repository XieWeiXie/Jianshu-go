package jianshu

type Publication struct {
}

func NewPublication() *Publication {
	return &Publication{}
}

func (p *Publication) GetPublicizedBook() {}

func (p *Publication) GetRecommendPublicizedAuthors() {}

func (p *Publication) GetNovelBooks() {}

func (p *Publication) GetITAndJobMarket() {}

func (p *Publication) GetCultureAndHistory() {}

func (p *Publication) GetMonthlyMagazine() {}
