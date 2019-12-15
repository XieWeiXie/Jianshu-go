package example

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/jianshu-go"
	"testing"
)

func TestExample(t *testing.T) {

	//fmt.Println("**** User ****")
	//user := jianshu.NewUser("https://www.jianshu.com/u/58f0817209aa", "", "")
	//fmt.Println("作者简书ID: ", user.GetUserID())
	//fmt.Println("作者简书主页: ", user.GetUserLink())
	//fmt.Println("作者性别: ", user.GetUserGender())
	//fmt.Println("作者关注数: ", user.GetFollowNumber())
	//fmt.Println("作者粉丝数: ", user.GetFollowerNumber())
	//fmt.Println("作者文章数: ", user.GetPassageNumber())
	//fmt.Println("作者写字数目: ", user.GetWriteNumber())
	//fmt.Println("作者收获喜欢数目: ", user.GetLikeNumber())
	//fmt.Println("作者主页文章信息: ", user.GetHomepagePassage())
	//fmt.Println("作者个人介绍: ", user.GetPersonalDetail())
	//fmt.Println("作者微博信息: ", user.GetTwitterInfo())
	//fmt.Sprint("作者喜欢的文章信息: ", user.GetLikedNotes())
	//fmt.Println("作者关注文集信息: ", user.GetSubscription())
	//fmt.Println("作者最新动态: ", user.GetLatestActive())
	//fmt.Println("作者最新评论: ", user.GetLatestCommented())
	//fmt.Println("作者热门文章: ", user.GetHotPassage())
	//
	//// fmt.Println("** Article **")
	//article := jianshu.NewArticle("https://www.jianshu.com/p/4ea9abf7e4e8")
	//fmt.Println("文章作者: ", article.GetAuthor())
	//fmt.Println("文章作者简介: ", article.GetDescription())
	//fmt.Println("文章标题: ", article.GetTitle())
	//fmt.Println("文章内容: ", article.GetContent())

	// fmt.Println(" ** Home-page **")
	//homePage := jianshu.NewHomePage(4, 1)
	//fmt.Println("** 获取首页文章 **")
	//fmt.Println(homePage.GetHomePagePassages())
	//fmt.Println("** 获取首页新上榜文章 **")
	//homePage.GetNewList()
	//fmt.Println("** 获取首页 7 日热门文章 **")
	//homePage.GetHotSeven()
	//fmt.Println("** 获取首页 30 日热门文章 **")
	//homePage.GetHotMonth()
	//fmt.Println("** 获取简书大学堂文章 **")
	//homePage.GetJianShuSchool()
	//fmt.Println("** Home-page-recommend **")
	//recommendAuthor := jianshu.NewRecommendAuthor(10)
	//fmt.Println("获取简书推荐10位作者: ", recommendAuthor.GetListRecommendAuthor())

	// fmt.Println("** Home-page-topic **")
	//topic := jianshu.NewTopic(12)
	//fmt.Println("获取推荐专题 12 个: ", topic.GetTopicCollectionRecommend())
	//fmt.Println("获取热门专题 12 个: ", topic.GetTopicCollectionHot())
	//fmt.Println("获取城市专题 12 个: ", topic.GetTopicCollectionCity())
	//fmt.Println("获取校园专题 12 个: ", topic.GetTopicCollectionSchoolyard())
	//fmt.Println("** Publication **")
	//publication := jianshu.NewPublication()
	//fmt.Println("获取简书已出版图书")
	//publication.GetPublicizedBook()
	//fmt.Println("获取简书小说")
	//publication.GetNovelBooks()
	//fmt.Println("获取简书IT、理财、职场图书")
	//publication.GetITAndJobMarket()
	//fmt.Println("获取简书文化、历史")
	//publication.GetCultureAndHistory()
	//fmt.Println("获取简书专题月刊")
	//publication.GetMonthlyMagazine()

	fmt.Println("** Special-subject **")
	specialSubject := jianshu.NewSpecialSubject("https://www.jianshu.com/c/dfcf1390085c")
	fmt.Println("获取专题标题: ", specialSubject.GetSpecialSubjectTitle())
	fmt.Println("获取专题公告: ", specialSubject.GetSpecialSubjectNotice())
	fmt.Println("获取专题文章数目: ", specialSubject.GetSpecialSubjectPassageNumber())
	fmt.Println("获取专题文章关注人数: ", specialSubject.GetSpecialSubjectFollowerNumber())
	fmt.Println("获取专题创建者信息: ")
	jianshu.JsonPretty(specialSubject.GetSpecialSubjectTitleAdministrator())
	fmt.Println("获取专题最新评论文章信息: ", specialSubject.GetSpecialSubjectNewComment())
	fmt.Println("获取专题最新收录文章信息: ", specialSubject.GetSpecialSubjectNewAdd())
	fmt.Println("获取专题最热门文章信息: ", specialSubject.GetSpecialSubjectHot())

}
