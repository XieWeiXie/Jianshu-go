package commands

import (
	"github.com/wuxiaoxiaoshen/jianshu-go"

	"fmt"

	"github.com/urfave/cli"
)

var newHomePageRecommend *jianshu.RecommendAuthor

func GetHomePageRecommendFlag() []cli.Flag {
	return nil
}

func getHomePageRecommendFlag() []cli.Flag {
	return []cli.Flag{
		cli.IntFlag{
			Name:  "number",
			Usage: "the recommend author to display",
		},
	}
}

func actionHomePageRecommend(c *cli.Context) {
	if c.NArg() < 1 {
		fmt.Println("execute command recommend method --number='*' <string>, should add one argument")
		return
	}
	if c.NArg() == 1 {
		startHomePageRecommend(c.Int("number"), c.Args()[0])
	}
}

func SubCMDHomePageRecommend() []cli.Command {
	return []cli.Command{
		{
			Name:   "method",
			Usage:  "the recommend author",
			Flags:  getHomePageRecommendFlag(),
			Action: actionHomePageRecommend,
		},
	}
}

func startHomePageRecommend(number int, method string) {
	newHomePageRecommend = jianshu.NewRecommendAuthor(number)
	if method == "get-list-recommend-author" || method == "1" {
		fmt.Println(newHomePageRecommend.GetListRecommendAuthor())
	}
}
