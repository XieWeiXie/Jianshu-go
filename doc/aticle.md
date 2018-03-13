## ARTICLE

### 接口:

- GetAuthor 获取文章作者
- GetDescription 获取作者简介
- GetTitle 获取文章标题
- GetContent 获取文章全文


### 示例:


```go

package main

import (
	"fmt"
	"jianshu-go"
)

func main(){
    article := jianshu.NewArticle("https://www.jianshu.com/p/4ea9abf7e4e8")
    fmt.Println("文章作者: ", article.GetAuthor())
    fmt.Println("文章作者简介: ", article.GetDescription())
    fmt.Println("文章标题: ", article.GetTitle())
    fmt.Println("文章内容: ", article.GetContent())
}

```

``` 

文章作者:  谢小路
文章作者简介:  上海大学2017级研究生毕业.微信公众号：Siwei_Jingjin
文章标题:  『简书API : jianshu 基于 golang （1）』
文章内容:  在我眼中，比较崇拜三类人：一类是设计师；一类是作家；一类是程序员。
这三类人都是通过创造、或者改善作品，不断的把世界变的更好。每每看到大师级的作品，总会不禁感叹，人与人的差别就是这么大。但是这都不阻碍我们模仿学习他们，向着更好的方向前进。
前几年，偏爱好于作家，总幻想自己能通过作品改变世界。后来证明，这条道路在真实的社会上，需要很大的毅力坚持，而且还需要点天分。
随着毕业、工作。我更偏爱设计师和程序员，而且两者在某些层面上有些共性。编程是我的本职工作，设计领域则是业余时间喜欢关注的点。
这三类人都在通过作品，不断的显现自己的能力。
所以一个程序员，假如没有开源作品，这样显的很格调不高。
开源作品质量其实也参差不齐。
一个好的开源作品：

代码质量优
解决的问题有实际用处
良好的维护
良好的文档

凡是都有第一步，第一步总是有各种各样的缺点，但这并不是不开源的理由。
也许吐槽的多了，或者别人给的意见多了。修改的多了，质量就更好了。

本项目尝试解析简书API。

编程语言：golang

主要的接口包括：

User: 个人主页信息
Article : 某篇文章的信息
Home-page: 简书主页的信息
Home-page-recommend: 简书推荐作者的信息
Home-page-topic: 简书推荐的专题信息
Publication： 简书出版信息
...


```