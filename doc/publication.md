## PUBLICATION

### 接口:



- GetPublicizedBook 获取已出版图书
- GetNovelBooks 获取小说
- GetITAndJobMarket 获取IT、理财、职场
- GetCultuereAndHistory 获取文化、历史
- GetMonthlyMagazine 获取专题月刊


### 示例:

```go

package main

import (
	"fmt"
	"jianshu-go"
)

func main(){
    publication := jianshu.NewPublication()
    fmt.Println("获取简书已出版图书")
    publication.GetPublicizedBook()
    fmt.Println("获取简书小说")
    publication.GetNovelBooks()
    fmt.Println("获取简书IT、理财、职场图书")
    publication.GetITAndJobMarket()
    fmt.Println("获取简书文化、历史")
    publication.GetCultureAndHistory()
    fmt.Println("获取简书专题月刊")
    publication.GetMonthlyMagazine()
}
```

```
获取简书已出版图书
[{陈信诚 转型王道:微商、社交电商实战系统 None ￥ 69.80 https://book.douban.com/subject/27622583/} {顾一宸 管他努力有没有回报，拼过才是人生 None ￥ 38.00 https://book.douban.com/subject/27596903/} {徐成东 从职场小白到团队老大：职场基本思维 None ￥ 39.00 https://book.douban.com/subject/27620324/} {简书 爱上写作，遇见更好的自己 None ￥ 38.00 https://book.douban.com/subject/27620180/} {李陌 失控的布局 出来混，总是要还的！但是，真相并非如此…… ￥ 35.00 https://book.douban.com/subject/27615008/ } {佰稼 别害怕，我们都孤寡 这七篇故事里的人孤独得难以自拔，因为现在就是一个孤独的时代。人与人之间情感越来越淡薄，正因为我们重感情，所以才不敢轻易地付以真情。 ￥ 39.00 https://book.douban.com/subject/27199582/} {尚元用 不要跟踪我 细腻入微的情感，环环相扣、跌宕起伏的情节，出人意料的结局，爱情与悬疑的完美结合。 ￥ 32.00 https://book.douban.com/subject/27113877/} {衷曲无闻 梦想不会辜负努力的你 写给对自己寄予厚望的你 ￥ 38.00 https://book.douban.com/subject/27077802/} {简书 总要为梦想疯狂一次 那些年陪你追梦的人还在身边吗？简书燃情作家和你同追梦共青春。 ￥ 36.00 https://book.douban.com/subject/27065409/} {简书 世界和它的悲欢 这世界永恒演绎着它的悲欢，故事里总能找到你我的影子。 ￥ 42.00 https://book.douban.com/subject/27008382/} {简书 终有人住进你心里 蛾扑火的勇气、奋不顾身的着迷、故作淡定的祝福… ￥ 36.00 https://book.douban.com/subject/26980212/} {沐丞 理财要趁早 理财不是发财，理财也不是简单的是投资、让钱生钱。 ￥ 30.00 https://book.douban.com/subject/26943302/} {简书 越勇敢的女人越幸运 本书是简书出品的第二本书，温暖又犀利地讲述了女人应该勇敢面对生活、爱情、自我成长等方面所遇到的问题和矛盾 ￥ 42.00 https://book.douban.com/subject/26926348/} {简书 你一定要努力，但千万别着急 唯有承担起厚重的经历，才能禁得起岁月的推敲。 ￥ 42.00 https://book.douban.com/subject/26786048/}]
获取简书小说
[{芝小麻 我离婚后的第一年 None ￥ 1.99 http://www.jianshu.com/p/11a941858272} {一元亦有用 神说要有光，我便开了灯 None 免费 http://www.jianshu.com/p/d2303a3b5b41} {离影疏落 读懂你的心——我的心理咨询师手记 None ￥ 1.99 http://www.jianshu.com/p/ab6a63674f62} {潮范 被骗入传销的13天 None ￥ 1.99 http://www.jianshu.com/p/b449e1fd14c9} {B杜 法兰西情人 None ￥ 12.00 http://www.jianshu.com/p/222d00822ea3} {无戒 我的山沟沟，我的女人 None ￥ 1.99 http://www.jianshu.com/p/832453f46086} {佰稼 再见，吾爱 None ￥ 1.99 http://www.jianshu.com/p/cfe52ce819ea} {简书短篇小说专题 你好，2027 None 免费 http://www.jianshu.com/p/855fdff634f0} {佰稼 无爱纪 None ￥ 1.99 http://www.jianshu.com/p/007fe8f6568e} {大肚子猫 爱在郁金香之国 None ￥ 3.99 http://www.jianshu.com/p/53754b3e7492}]
获取简书IT、理财、职场图书
[{齐帆齐 我的草根奋斗故事 None ￥ 1.99 http://www.jianshu.com/p/b2e0a9e9b804} {徐成东 《东哥说职场》第二辑 高效沟通与情绪管理 None ￥ 2.99 http://www.jianshu.com/p/2b7b235492b5} {沐丞 买房要趁早 None ￥ 9.99 http://www.jianshu.com/p/be9de8b0c845} {徐成东 《东哥说职场》第一辑 求职与加薪的正确姿势 None ￥ 4.99 http://www.jianshu.com/p/2e591ca7435b} {蓝桥飞 游戏策划入门修行 None ￥ 1.99 http://www.jianshu.com/p/80e630e81c2a} {庄表伟 开源思索集 None ￥ 1.99 http://www.jianshu.com/p/e5c5af1e45dc} {沐丞 工作十年买下三套房 None ￥ 1.99 http://www.jianshu.com/p/bf306d2725f2} {纯银V 10个案例说明什么是产品模型 None ￥ 1.99 http://www.jianshu.com/p/f74527bd0e04} {bylin 取悦的工序：如何理解游戏 None 免费 http://www.jianshu.com/p/d4be65e0b792} {萌大夫精神病院 我也曾经上班一个礼拜就辞职 None 免费 http://www.jianshu.com/p/d99d0cc4575e}]
获取简书文化、历史
[{芊语芊寻 当你想要学民乐》 None ￥ 3.99 http://www.jianshu.com/p/65c3c8163162} {陈缃眠 故人之书 None ￥ 1.99 http://www.jianshu.com/p/55f735010b37} {王祎 没想到你是这样的美术史 None ￥ 2.99 http://www.jianshu.com/p/d706466fdee3} {马风 宋词·萧红·林徽因/一丘一壑也风流 None ￥ 6.99 http://www.jianshu.com/p/ae719ebb0cbc} {一个历史围观群众 民国人物趣味杂谈 None ￥ 1.99 http://www.jianshu.com/p/c1f68cf064c0} {饱醉豚 新加坡的那些事儿 None ￥ 2.99 http://www.jianshu.com/p/a11911c81ed0}]
获取简书专题月刊
[{简书茶馆叶老板 简书茶点故事精选第二辑 None 免费 http://www.jianshu.com/p/156896ff347e} {简书教育专题 简书教育月刊002 · 当我们谈论教育时，我们谈论的是什么 None 免费 http://www.jianshu.com/p/a368a16c75fe} {简书短篇小说专题 简书短篇小说月刊005 · 传奇故事 世间冷暖 None 免费 http://www.jianshu.com/p/6472808c21bd} {简书谈写作专题 简书谈写作月刊003 · 品读岁月，书写人生 None 免费 http://www.jianshu.com/p/dc8134000909} {简书读书专题 简书读书月刊001 · 穷理尽性 以至于命 None 免费 http://www.jianshu.com/p/78d0b28fe8f8} {简书世间事专题 简书世间事月刊001 · 不动声色 None 免费 http://www.jianshu.com/p/ae6bc1b682d0} {简书奇思妙想专题 简书奇思妙想月刊001 · 有梦为马，随处可栖 None 免费 http://www.jianshu.com/p/1d3f71fadc63} {简书茶馆叶老板 简书茶点故事7月刊 None 免费 http://www.jianshu.com/p/56217ea0170a} {简书谈写作专题 简书谈写作月刊002 · 文字情缘 None 免费 http://www.jianshu.com/p/cf12eedb4536} {简书人物专题 简书人物月刊001 · 再回首绝代芳华 None 免费 http://www.jianshu.com/p/f1dd98c85e13}]

```