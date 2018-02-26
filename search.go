package jianshu

type HomePageSearch struct {
	query string
}

func NewSearch(query string) *HomePageSearch {
	return &HomePageSearch{
		query: query,
	}
}

func (s *HomePageSearch) GetAuthor(number int) {
	// number < 3
}

func (s *HomePageSearch) GetPassage(number int) {
	// number < 9
}

func (s *HomePageSearch) GetAllAuthorNumber() {}

func (s *HomePageSearch) GetAllPassageNumber() {}
