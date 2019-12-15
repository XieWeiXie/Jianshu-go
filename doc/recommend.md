## HOME-PAGE-RECOMMEND


### 接口：

- GetListRecommendAuthor 获取首页推荐作者



### 示例:


```go


package main

import (
	"fmt"
	"github.com/wuxiaoxiaoshen/jianshu-go"
)

func main(){
    recommendAuthor := jianshu.NewRecommendAuthor(10)
    fmt.Println("获取简书推荐10位作者: ", recommendAuthor.GetListRecommendAuthor())
}

```

``` 
获取简书推荐10位作者:  [{寻麦 https://www.jianshu.com/users/e0ef486d9b90 None 卖童话的小巫师，强迫症晚期。 《菜菜的奇幻之旅》第十二章：木匠的茶话会致读者：在妖怪的世界里，找回丢失的童真和勇气《菜菜的奇幻之旅》第十一章：自言自语的短鼻象} {遛遛心情的溜妈 https://www.jianshu.com/users/dbfdce352c0d None 溜爸，简书签约作者，舞枪弄棒、舞文弄墨的计算... 清明，约你一起种田，读书……一本吸引我前前后后读了三遍的书，和你分享（下）［奇幻］唐朝那些猫事儿（57）} {驴光掠影 https://www.jianshu.com/users/b3b2c03354f3 None 曾在千亿级A股上市公司任IT高管，现任某创业... 认知系列|知识留存率最高的阅读原则随笔|我去年买了个币云南滇西|你去了一个假的大理} {简书出版 https://www.jianshu.com/users/55b597320c4e None 给我发简信前请先看完我的简介。关于如... 简书2017年的10个好故事（虚构&非虚构）上线整整十部优质付费连载，让你周末看个够更新一小波实用性非常强的优质付费连载} {末晓 https://www.jianshu.com/users/d0f9f963aeb9 None 每一个故事都是写给自己，希望路过的你们恰好喜欢 ［童话］花仙卡卡的魔法修行（7）［童话］花仙卡卡的魔法修行（6）花仙卡卡的魔法修行（5）} {雨落荒原 https://www.jianshu.com/users/7406f22f461e woman 反反香菜联盟会会长合作请联系简书版权... 【3000字转折大赛】两枚硬币《对岸》（10）被男人包围的女人-往事关于写作的一些个人体验} {格列柯南 https://www.jianshu.com/users/ffc565d738a3 man 历史博士、哲学硕士，本科经济学。只有审视生活... 四十而立，五十不惑，六十用伟哥《三言二拍》教你如何写豹纹父母反应过度与失败的孩子教育} {简书大学堂 https://www.jianshu.com/users/c5580cc1c3f4 None 简书自有的学习成长平台。简书大学堂讲... 被忽视的植树节，学堂菌决定来种草上架这个功能，“累死"了一位产品经理这24条职场弯路，你本可以绕开的...} {aloho https://www.jianshu.com/users/1446a350e58a man aloho。名字改编自夏威夷语aloha，你... 有种稳赚不赔的买卖叫环游中国|《绝对光年》番外篇aloho的床头诗601-650一道数学题，解答你能不能环游中国的困惑|《绝对光年》番外篇} {魏童 https://www.jianshu.com/users/5462ec6828f6 man 魏童，知名图书策划人、影视策划人。先后在盛大... 如何发现你的创作天赋和优势？天才的创作者与精明的经营者世间的那些相遇}]


```