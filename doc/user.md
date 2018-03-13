## USER

### 接口:

- GetUserID  获取用户ID
- GetUserLink 获取用户主页URL
- GetUserGender 获取用户性别
- GetFollowNumber 获取用户关注数
- GetFollowerNumber 获取用户粉丝数
- GetPassageNumber 获取用户文章书面
- GetWriteNumber 获取用户写的字数
- GetLikeNumber 获取用户得到的喜欢的数目
- GetHomePagePassage 获取用户主页文章信息
- GetPersonalDetail 获取用户个人介绍
- GetTwitterInfo 获取用户微博地址
- GetLikedNotes 获取作者喜欢的文章信息
- GetSubscription 获取作者关注的专题/文集/连载
- GetLatestActice 获取作者最新动态
- GetLatestCommented 获取作者最新评论
- GetHotPassage 获取作者热门文章


### 示例：

```go

package main

import (
	"fmt"
	"jianshu-go"
)

func main() {
    user := jianshu.NewUser("https://www.jianshu.com/u/58f0817209aa", "", "")
    fmt.Println("作者简书ID: ", user.GetUserID())
    fmt.Println("作者简书主页: ", user.GetUserLink())
    fmt.Println("作者性别: ", user.GetUserGender())
    fmt.Println("作者关注数: ", user.GetFollowNumber())
    fmt.Println("作者粉丝数: ", user.GetFollowerNumber())
    fmt.Println("作者文章数: ", user.GetPassageNumber())
    fmt.Println("作者写字数目: ", user.GetWriteNumber())
    fmt.Println("作者收获喜欢数目: ", user.GetLikeNumber())
    fmt.Println("作者主页文章信息: ", user.GetHomepagePassage())
    fmt.Println("作者个人介绍: ", user.GetPersonalDetail())
    fmt.Println("作者微博信息: ", user.GetTwitterInfo())
    fmt.Sprint("作者喜欢的文章信息: ", user.GetLikedNotes())
    fmt.Println("作者关注文集信息: ", user.GetSubscription())
    fmt.Println("作者最新动态: ", user.GetLatestActive())
    fmt.Println("作者最新评论: ", user.GetLatestCommented())
    fmt.Println("作者热门文章: ", user.GetHotPassage())
	
}

```

``` 
作者简书ID:  谢小路
作者简书主页:  https://www.jianshu.com/u/58f0817209aa
作者性别:  None
作者关注数:  9
作者粉丝数:  562
作者文章数:  83
作者写字数目:  85308
作者收获喜欢数目:  470
作者主页文章信息:  [{2018-03-11T20:12:04+08:00 『简书API : jianshu 基于 golang （1）』 在我眼中，比较崇拜三类人：一类是设计师；一类是作家；一类是程序员。 这三类人都是通过创造、或者改善作品，不断的把世界变的更好。每每看到大师级的作品，总会不禁感叹，人与人的差别... 100 0 2 1 谢小路} {2018-03-10T23:17:32+08:00 『Python 爬虫文集梳理』 过去的几年内，我开始了编程。 过去的一年内，我开始了工作生涯。 我学会的第一个编程技能是『爬虫』，工作后，开始接触Golang。 我开始不断的将编程结合业务， 接触越来越多的... 26 0 2 0 谢小路} {2018-03-10T00:23:23+08:00 『Ansible 上手指南：2』 读一本书最好的时机是什么时候？是你刚买的时候，趁着新鲜劲，先了解这本书，继而马上阅读完这本书。如果错过了最好的时机阅读一本书，那什么时候是合适的时机，是你需要这方面的资料或者... 7 2 1 0 谢小路} {2018-03-06T00:13:44+08:00 『requests-html 源码学习:  1』 大家好，我是谢伟，是一名程序员，熟悉 Pyhton 和 Go。学会的第一个技能是『网络爬虫』。 最近 Python 领域大神  kennethreitz 开源了一个关于网络内... 30 0 3 0 谢小路} {2018-03-05T21:37:43+08:00 『Ansible 上手指南』 Ansible 上手指南 前言 最近在重构一款命令行工具，使用 golang 重新开发，但需要继续维持原有的命令，同时增加新命令。 在重构的过程中，需要对现命令行工具和原命令... 24 0 2 0 谢小路} {2018-02-11T23:03:25+08:00 『Go 语言实现简易爬虫：市值前100数字货币交易信息』 大家好，我是谢伟，一名程序员。之前接触的语言是Python， 编程领域学会的第一个技能是『爬虫』，凭借着爬虫技术先后在两个创业公司从事的是『网络爬虫』这份活。 研究生毕业后，... 80 0 8 0 谢小路} {2018-02-08T22:56:47+08:00 『Beego + Swagger 快速上手』 大纲 Beego 是什么 为什么写这个 如何指导 前几天我写了一个Swagger 上手指南，觉得还是让使用者难以上手。尽管它是一款优秀的API 工具。 但我在编写API 的过... 72 0 0 0 谢小路} {2018-02-03T15:23:06+08:00 『2018年1月知识点合集』 我有一个习惯，就是不断的记录在工作中反复用到的知识点，原本我很喜欢使用印象笔记和有道云笔记，其一是云笔记的同步功能，其二是云笔记的搜索功能，当你输入的笔记多了之后，你才会发现... 100 0 0 0 谢小路} {2018-01-30T21:57:42+08:00 『Swagger 上手』 大纲 问题 RestfulAPI API 动作 请求：Url、Body 返回信息：Status_code、Response 在开发过程中，经常会遇到和其他组件或者服务进行交互... 39 0 1 0 谢小路}]
作者个人介绍:  上海大学2017级研究生毕业.微信公众号：Siwei_Jingjin
作者微博信息:  http://weibo.com/u/1948244870
作者关注文集信息:  [{微信小程序开发于连林520wcf https://www.jianshu.com/c/dfdc2bbd1315 于连林520wcf 收录了660篇文章 3037人关注} {PPT：千页计划谢小路 https://www.jianshu.com/c/d76906de84c5 谢小路 收录了1篇文章 1人关注} {Go 语言学习专栏 谢小路 https://www.jianshu.com/c/0d857ad2e739 谢小路 收录了8篇文章 1人关注} {PPT沈晓马 https://www.jianshu.com/c/9df5862d2a58 沈晓马 收录了1603篇文章 8320人关注} {我的Python自学之路帅灰 https://www.jianshu.com/c/0690e20b7e7d 帅灰 收录了1047篇文章 5183人关注} {强化训练：python谢小路 https://www.jianshu.com/c/03322437b28d 谢小路 收录了19篇文章 20人关注} {爬虫专栏谢小路 https://www.jianshu.com/c/dfcf1390085c 谢小路 收录了262篇文章 950人关注}]
作者最新动态:  [{comment_note 『Ansible 上手指南：2』} {share_note 『简书API : jianshu 基于 golang （1）』} {share_note 『Python 爬虫文集梳理』} {like_user 若与} {share_note 『Ansible 上手指南：2』} {like_user xiaopiu原型设计} {share_note 『requests-html 源码学习:  1』} {share_note 『Ansible 上手指南』} {like_collection } {like_user 刘英滕} {share_note 『Go 语言实现简易爬虫：市值前100数字货币交易信息』} {comment_note 我儿子} {comment_note 修正我的学习} {share_note 『Beego + Swagger 快速上手』} {like_note 【9页幻灯片】阿尔卑斯山} {share_note 『2018年1月知识点合集』} {share_note 『Swagger 上手』} {share_note 『Go 语言学习专栏』-- 第一期} {comment_note 2018『PPT 千页计划』第一期} {share_note 2018『PPT 千页计划』第一期}]
作者最新评论:  [{『Ansible 上手指南：2』 读一本书最好的时机是什么时候？是你刚买的时候，趁着新鲜劲，先了解这本书，继而马上阅读完这本书。如果错过了最好的时机阅读一本书，那什么时候是合适的时机，是你需要这方面的资料或者... https://www.jianshu.com/p/82245bb31bcd} {2018『PPT 千页计划』第一期 2018 年我准备花上一年的业余时间制作 1000 页PPT。 没什么目的。 就是制作 1000 页 PPT。 内容来自一些作业，也来自一些我喜欢作者的原文模仿，也来自一些模... https://www.jianshu.com/p/492875661ab8} {如何成为一名设计师.ppt 我是一名程序员，业余时间玩PPT。 这是根据罗子雄的 TED 演讲而制作的PPT。 很费时间，但习得一门技能从来都不是那么轻易和容易。 https://www.jianshu.com/p/88e861223299} {线性代数：一切为了更好的理解 ![][1][1]: http://latex.codecogs.com/gif.latex?\alpha 线性代数是数学工具掌握它，打开数学的另一扇大门 1：声明 非原创，... https://www.jianshu.com/p/c0e119c76971} {技术文档如何编写？ 关于文档编写的几个思维 近期重新组织了好几篇技术文档，把其中的一些感悟提炼出来。 文档为达到容易理解和操作的程度，对大量的语言重新组织，内容的不同呈现，借助辅助工具等一系列操... https://www.jianshu.com/p/6d66262f6dcd} { 001     |    工作：一周总结 已毕业，已入职，本职工作是云平台部署和管理， PAAS层。 刚入手，系列一周总结，会记录一些工作的学习和思考，目标是：提高业务能力和专业技能，同时成为一名合格的职场人。 为实... https://www.jianshu.com/p/b9eedf83675c} {Python: 一周笔记 主要介绍几个用到的python模块的使用方法。python 含有丰富的内置和第三方库，企图全部掌握并精通那是不可能的。 但当你开发任务需要到的时候，你可以及时的避免重复的一些... https://www.jianshu.com/p/ad450b2126b0} {002 | 工作：职场学习术 这篇文章的灵感来源于暂停玩「王者荣耀」之后，甚至我产生了卸载的冲动，我觉得上班的精神状态和这款手游有着潜移默化的关系。 好，撇开这些不谈，我想谈谈在职场中思考的比较多的两件事... https://www.jianshu.com/p/03c0849eced8} {004 | 工作：月度总结 时间过的很快，距离毕业离校已经过去将近三个月，上班也一个快一个半月，每天都在不断的接触新的东西，新旧知识需要融合，每周公司新人需要进行周报总结，算是一次新旧知识融合，但是输出... https://www.jianshu.com/p/cdb22fa377d1}]
作者热门文章:  [{2017-01-15T09:56:51+08:00 Python:一周笔记 主题 邮件处理 日志模块 pdf处理 md5 mongodb索引和聚合 excel 读写 1. 发送邮件模块 这里指的邮件功能当然不是指的是职场上所谓的邮件，指的是程序运行中... 939 3 30 0 谢小路} {2016-04-29T18:36:33+08:00 专栏：005：Beautiful Soup 的使用 系列爬虫专栏 崇尚的学习思维是：输入，输出平衡，且平衡点不断攀升。 曾经有大神告诫说：没事别瞎写文章；所以，很认真的写的是能力范围内的，看客要是看不懂，不是你的问题，问题在我... 600 2 27 0 谢小路} {2017-08-06T15:23:39+08:00 Python: 一周笔记 主要介绍几个用到的python模块的使用方法。python 含有丰富的内置和第三方库，企图全部掌握并精通那是不可能的。 但当你开发任务需要到的时候，你可以及时的避免重复的一些... 580 3 22 0 谢小路} {2016-04-30T12:23:51+08:00 专栏：006：实战爬取博客 （起什么优雅的名字点击量会增多？） 系列爬虫专栏 崇尚的学习思维是：输入，输出平衡，且平衡点不断攀升。 曾经有大神告诫说：没事别瞎写文章；所以，很认真的写的是能力范围内的，看... 696 1 21 0 谢小路} {2016-05-13T21:15:47+08:00 番外篇：面试总结(1) 初学者 你经历的每一件事都会成为未来的部分，具体看你如何对待了. 0：前言 作为一个初学者，对知识的理解存在着很多的疑惑。同人交流作为学习的方式之一，牛人和兴趣的着眼点的不同... 925 10 18 0 谢小路} {2016-04-27T17:59:40+08:00 专栏：003：正则表达式 系列爬虫专栏 崇尚的学习思维是：输入，输出平衡，且维持平衡点不断精进的地步 曾经有大神告诫说：没事别瞎写文章；为此写的都是，在我能力范围内的 1：框架 2：概念 什么是正则表... 547 5 15 0 谢小路} {2016-05-21T11:43:06+08:00  专栏：016：功能强大的“图片下载器” 用理工科思维看待这个世界 系列爬虫专栏 初学者，尽力实现最小化学习系统 如何实现项目图片的下载 0：学习理念 推荐阅读简书：学习方法论我觉得对我有帮助，多问自己为什么从来不是... 1269 2 14 0 谢小路} {2016-06-02T15:02:33+08:00 线性代数：一切为了更好的理解 ![][1][1]: http://latex.codecogs.com/gif.latex?\alpha 线性代数是数学工具掌握它，打开数学的另一扇大门 1：声明 非原创，... 943 4 14 0 谢小路} {2016-11-13T13:46:36+08:00  Python 强化训练：第十篇 主题： pycharm 开发利器，减少编写错误代码的可能性，遵从PEP8编码规范. 配置 设置 安装模块 快捷键 Git 版本控制 1. 配置 setting--> proj... 437 3 14 0 谢小路}]


```