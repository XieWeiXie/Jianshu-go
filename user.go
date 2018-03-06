package jianshu

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type User struct {
	userID string
	gender string
	Link   string
	Pages  int
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

func (user *User) doc() *goquery.Document {
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
	doc := user.doc()
	name := doc.Find("div.title").Find("a.name").Text()
	user.userID = name
	return name
}

func (user *User) GetUserLink() string {
	if user.Link != "" {
		return user.Link
	}
	doc := user.doc()
	tempLink, _ := doc.Find("div.title").Find("a.name").Attr("href")
	link := MakeCompleteUrl(tempLink)
	return link
}

func (user *User) GetUserGender() string {
	if user.gender != "" {
		return user.gender
	}
	doc := user.doc()
	tempGender, err := doc.Find("div.title").Find("i").Attr("class")
	if !err {
		return "None"
	}
	user.gender = SplitString(tempGender)
	return SplitString(tempGender)

}

func (user *User) GetFollowNumber() int {
	doc := user.doc()
	tempFollowNumber := doc.Find("div.info ul li").Eq(0).Find("div a p").Text()
	return StringToInt(tempFollowNumber)
}

func (user *User) GetFollowerNumber() int {
	doc := user.doc()
	tempFollowerNumber := doc.Find("div.info ul li").Eq(1).Find("div a p").Text()
	return StringToInt(tempFollowerNumber)
}

func (user *User) GetPassageNumber() int {
	doc := user.doc()
	tempPassageNumber := doc.Find("div.info ul li").Eq(2).Find("div a p").Text()
	user.Pages = Pages(StringToInt(tempPassageNumber))
	return StringToInt(tempPassageNumber)
}

func (user *User) GetWriteNumber() int {
	doc := user.doc()
	tempWriteNumber := doc.Find("div.info ul li").Eq(3).Find("div p").Text()
	return StringToInt(tempWriteNumber)
}

func (user *User) GetLikeNumber() int {
	doc := user.doc()
	tempLikeNumber := doc.Find("div.info ul li").Eq(4).Find("div p").Text()
	return StringToInt(tempLikeNumber)
}

// 获取个人主页文章集合：最多9篇
func (user *User) GetHomepagePassage() []PassageDetail {
	var allPassage []PassageDetail
	doc := user.doc()
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
	doc := user.doc()
	tempDescription := StringSpace(doc.Find("div.description div.js-intro").Text())
	return tempDescription
}

func (user *User) GetTwitterInfo() string {
	var twitter string
	doc := user.doc()
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

func (user *User) GetLikedNotes() []LikeNotes {
	var allLikedPassages []LikeNotes
	doc, _ := goquery.NewDocument(LikedNotesUrl(user))
	likedNumber := doc.Find("ul.trigger-menu li").Eq(1).Find("a").Text()
	if StringGetInt(likedNumber) == 0 {
		return nil
	}
	doc.Find("div.content").Each(func(i int, selection *goquery.Selection) {
		if i >= 1 {
			author := selection.Find("div.info a.nickname").Text()
			tempAuthorUrl, _ := selection.Find("div.info a").Attr("href")
			authorUrl := MakeCompleteUrl(tempAuthorUrl)
			title := selection.Find("div.info a.title").Text()
			abstract := StringSpace(selection.Find("p.abstract").Text())
			//fmt.Println(author, authorUrl, title, abstract)
			one := LikeNotes{
				author:    author,
				authorUrl: authorUrl,
				title:     title,
				abstract:  abstract,
			}
			allLikedPassages = append(allLikedPassages, one)
		}
	})
	return allLikedPassages
}

func (user *User) GetSubscription() []SpecialTopics {
	var allSubscription []SpecialTopics
	doc, _ := goquery.NewDocument(SubscriptionUrl(user))
	SubscriptionNumber := doc.Find("ul.trigger-menu li").Eq(0).Find("a").Text()
	if StringGetInt(SubscriptionNumber) == 0 {
		return nil
	}
	doc.Find("div.info").Each(func(i int, selection *goquery.Selection) {
		if i >= 1 {
			name := selection.Find("a").Text()
			tempUrl, _ := selection.Find("a").Attr("href")
			url := MakeCompleteUrl(tempUrl)
			creatorPassageAndFollower := StringSpace(selection.Find("div.meta").Text())
			creator, PassageAndFollower := StringHandle(creatorPassageAndFollower)
			passage, follower := StringSplitWith(PassageAndFollower)
			//fmt.Println(name, url, creator, passage, follower)
			one := SpecialTopics{
				name:                name,
				url:                 url,
				creator:             creator,
				totalPassage:        passage,
				totalFollowerNumber: follower,
			}
			allSubscription = append(allSubscription, one)
		}

	})
	return allSubscription
}

func (user *User) GetLatestActive() []LatestActive {
	var allLatestActive []LatestActive
	doc, _ := goquery.NewDocument(TimeLineUrl(user))
	doc.Find("ul.note-list li").Each(func(i int, selection *goquery.Selection) {
		var activeDetail string
		tempActiveType, _ := selection.Find("span").Attr("data-type")
		activeType := StringSpace(tempActiveType)
		if strings.Contains(activeType, "like_comment") {
			activeDetail = StringSpace(selection.Find("p.comment").Text())
		}
		if strings.Contains(activeType, "like_user") {
			activeDetail = StringSpace(selection.Find("a.title").Text())
		}
		if strings.Contains(activeType, "comment_note") {
			activeDetail = StringSpace(selection.Find("a.title").Text())
		}
		if strings.Contains(activeType, "like_note") {
			activeDetail = StringSpace(selection.Find("a.title").Text())
		}
		if strings.Contains(activeType, "share_note") {
			activeDetail = StringSpace(selection.Find("a.title").Text())
		}
		//fmt.Println(activeType, activeDetail)
		one := LatestActive{
			activeType:   activeType,
			activeDetail: activeDetail,
		}
		allLatestActive = append(allLatestActive, one)
	})
	return allLatestActive
}

func (user *User) GetLatestCommented() []LatestCommented {
	var allLatestCommented []LatestCommented
	if user.GetWriteNumber() == 0 {
		return nil
	}
	doc, _ := goquery.NewDocument(CommentedUrl(user))
	doc.Find("ul.note-list li").Each(func(i int, selection *goquery.Selection) {
		title := selection.Find("a.title").Text()
		tempUrl, _ := selection.Find("a").Attr("href")
		url := MakeCompleteUrl(tempUrl)
		abstract := StringSpace(selection.Find("p.abstract").Text())
		//fmt.Println(title, url, abstract)
		one := LatestCommented{
			title:    title,
			url:      url,
			abstract: abstract,
		}
		allLatestCommented = append(allLatestCommented, one)
	})
	return allLatestCommented
}

func (user *User) GetHotPassage() []PassageDetail {
	var allPassageDetail []PassageDetail
	if user.GetWriteNumber() == 0 {
		return nil
	}
	doc, _ := goquery.NewDocument(HotPassageUrl(user))
	doc.Find("ul.note-list li").Each(func(i int, selection *goquery.Selection) {
		time, _ := selection.Find("div.content").Find("div.author").Find("div.info span").Attr("data-shared-at")
		title := selection.Find("a.title").Text()
		abstract := StringSpace(selection.Find("p.abstract").Text())
		reader := StringToInt(StringSpace(selection.Find("div.meta").Find("a").Eq(0).Text()))
		comment := StringToInt(StringSpace(selection.Find("div.meta").Find("a").Eq(1).Text()))
		liked := StringToInt(StringSpace(selection.Find("div.meta").Find("span").Eq(0).Text()))
		payed := StringToInt(StringSpace(selection.Find("div.meta").Find("span").Eq(1).Text()))
		author := user.userID
		//fmt.Println(time, title, abstract, reader, comment, liked, payed, author)
		one := PassageDetail{
			time:     time,
			title:    title,
			abstract: abstract,
			reader:   reader,
			comment:  comment,
			liked:    liked,
			payed:    payed,
			author:   author,
		}
		allPassageDetail = append(allPassageDetail, one)
	})
	return allPassageDetail
}
