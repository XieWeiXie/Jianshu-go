package jianshu

type HomePage struct {
	number int
	page   int
}

// number <= 20 page <= 10
func NewHomePage(number, page int) *HomePage {
	return &HomePage{
		number: number,
		page:   page,
	}
}

func GetHomePageTopics() {}

func GethomePagePassages() {}

func GetNewList() {}

func GetHotSeven() {}

func GetHotMonth() {}

func GetJianshuSchool() {}
