package jianshu

import (
	"strings"

	"fmt"

	"github.com/PuerkitoBio/goquery"
)

type User struct {
	userID string `json:"user_id"`
	gender string `json:"gender"`
	Link   string `json:"link"`
	Pages  int    `json:"pages"`
}

func NewUser(link string, gender string, userID string) *User {
	if link == "" {
		panic("The link should not be nil.")
	}
	return &User{
		Link:   link,
		userID: userID,
		gender: gender,
	}
}

func (user *User) Doc() *goquery.Document {
	if user.Link == "" {
		panic("The link should not be nil.")
	}
	doc, _ := goquery.NewDocument(user.Link)
	return doc
}

func (user *User) GetUserID() string {
	if user.userID != "" {
		return user.userID
	}
	doc := user.Doc()
	name := doc.Find("div.title").Find("a.name").Text()
	user.userID = name
	return name
}

func (user *User) GetUserLink() string {
	if user.Link != "" {
		return user.Link
	}
	doc := user.Doc()
	tempLink, _ := doc.Find("div.title").Find("a.name").Attr("href")
	link := MakeCompleteUrl(tempLink)
	return link
}

func (user *User) GetUserGender() string {
	if user.gender != "" {
		return user.gender
	}
	doc := user.Doc()
	tempGender, err := doc.Find("div.title").Find("i").Attr("class")
	if !err {
		return "None"
	}
	user.gender = SplitString(tempGender)
	return SplitString(tempGender)

}

func (user *User) GetFollowNumber() int {
	doc := user.Doc()
	tempFollowNumber := doc.Find("div.info ul li").Eq(0).Find("div a p").Text()
	return StringToInt(tempFollowNumber)
}

func (user *User) GetFollowerNumber() int {
	doc := user.Doc()
	tempFollowerNumber := doc.Find("div.info ul li").Eq(1).Find("div a p").Text()
	return StringToInt(tempFollowerNumber)
}

func (user *User) GetPassageNumber() int {
	doc := user.Doc()
	tempPassageNumber := doc.Find("div.info ul li").Eq(2).Find("div a p").Text()
	user.Pages = Pages(StringToInt(tempPassageNumber))
	return StringToInt(tempPassageNumber)
}

func (user *User) GetWriteNumber() int {
	doc := user.Doc()
	tempWriteNumber := doc.Find("div.info ul li").Eq(3).Find("div p").Text()
	return StringToInt(tempWriteNumber)
}

func (user *User) GetLikeNumber() int {
	doc := user.Doc()
	tempLikeNumber := doc.Find("div.info ul li").Eq(4).Find("div p").Text()
	return StringToInt(tempLikeNumber)
}

// 获取个人主页文章集合：最多9篇
func (user *User) GetHomepagePassage() []PassageDetail {
	var allPassage []PassageDetail
	doc := user.Doc()
	if user.GetWriteNumber() == 0 {
		return nil
	}
	if doc.Find("div#outer-container").Size() == 0 {
		doc.Find("ul.note-list li").Each(func(i int, selector *goquery.Selection) {
			var onePassage PassageDetail
			time, exist := selector.Find("div.content").Find("div.author").Find("div.info span").Attr("data-shared-at")
			if !exist {
				time = ""
			}
			title := selector.Find("div.content").Find("a.title").Text()
			abstract := StringSpace(selector.Find("div.content").Find("p").Text())
			author := StringSpace(selector.Find("div.content").Find("div.author div.info").Text())
			reader := StringToInt(StringSpace(selector.Find("div.content div.meta").Find("a").Eq(0).Text()))
			comment := StringToInt(StringSpace(selector.Find("div.content div.meta").Find("a").Eq(1).Text()))
			liked := StringToInt(StringSpace(selector.Find("div.content div.meta").Find("span").Eq(0).Text()))
			payed := StringToInt(StringSpace(selector.Find("div.content div.meta").Find("span").Eq(1).Text()))
			onePassage.title = title
			onePassage.abstract = abstract
			onePassage.author = author
			onePassage.reader = reader
			onePassage.comment = comment
			onePassage.liked = liked
			onePassage.payed = payed
			onePassage.time = time
			allPassage = append(allPassage, onePassage)
		})
	} else {
		onePassage := PassageDetail{}
		allPassage = append(allPassage, onePassage)
	}
	return allPassage
}

func (user *User) GetPersonalDetail() string {
	doc := user.Doc()
	tempDescription := StringSpace(doc.Find("div.description div.js-intro").Text())
	return tempDescription
}

func (user *User) GetTwitterInfo() string {
	var twitter string
	doc := user.Doc()
	tempInfoSize := doc.Find("div.description a").Size()
	if tempInfoSize == 0 {
		return "None"
	}
	doc.Find("div.description a").EachWithBreak(func(i int, selection *goquery.Selection) bool {
		info, exist := selection.Attr("class")
		if !exist {
			return false
		}
		if strings.Contains(info, "weibo") {
			twitter, _ = selection.Attr("href")
		}
		return true
	})
	return twitter
}

// 废弃
func (user *User) GetWeChatInfo() string {
	return ""
}

// 暂时无法解析url
func (user *User) GetTopics() []*Topics {
	var allTopics []*Topics
	return allTopics
}

// 暂时无法解析url
func (user *User) GetManagement() []*Topics {
	var allManagement []*Topics
	return allManagement
}

// 暂时无法解析url
func (user *User) GetPassageTopics() []*Topics {
	var allPassageTopics []*Topics
	return allPassageTopics
}

func (user *User) GetLikedNotes() []PassageDetail {
	var allLikedPassages []PassageDetail
	doc, _ := goquery.NewDocument(LikedNotesUrl(user))
	likedNumber := doc.Find("ul.trigger-menu li").Eq(1).Find("a").Text()
	if StringGetInt(likedNumber) == 0 {
		return nil
	}

	return allLikedPassages
}

func (user *User) GetSubscription() []SpecialTopics {
	var allLikedTopics []SpecialTopics
	doc, _ := goquery.NewDocument(SubscriptionUrl(user))
	SubscriptionNumber := doc.Find("ul.trigger-menu li").Eq(0).Find("a").Text()
	if StringGetInt(SubscriptionNumber) == 0 {
		return nil
	}
	doc.Find("div.info").Each(func(i int, selection *goquery.Selection) {
		//name := selection.Find("a.title").Text()
		//tempUrl, _ := selection.Find("a").Attr("href")
		//url := MakeCompleteUrl(tempUrl)
		creatorPassageAndFollower := selection.Find("div.meta").Text()
		creator := strings.Replace(creatorPassageAndFollower, "\n", "", -1)
		fmt.Println(creator)
	})
	return allLikedTopics
}

func (user *User) GetLatestInfo() *LatestComment {
	return &LatestComment{}
}

func (user *User) GetLatestContent() *PassageDetail {
	return &PassageDetail{}
}

func (user *User) GetHotPassage() *PassageDetail {
	return &PassageDetail{}
}
