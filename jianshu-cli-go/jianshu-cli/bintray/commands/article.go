package commands

import (
	"github.com/wuxiaoxiaoshen/jianshu-go"

	"fmt"

	"github.com/urfave/cli"
)

var newArticle *jianshu.Article

func GetArticleFlag() []cli.Flag {
	return nil
}

func getArticleFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "url",
			Usage: "get article passage url",
		},
	}
}

func actionArticle(c *cli.Context) {
	if c.NArg() < 1 {
		fmt.Println("execute command article  method --url='*****' <string>, should add one argument ")
		return
	}
	if c.NArg() == 1 {
		startArticle(c.String("url"), c.Args()[0])
	}
}

func SubCMDArticle() []cli.Command {
	return []cli.Command{
		{
			Name:   "method",
			Usage:  "get Article method",
			Flags:  getArticleFlag(),
			Action: actionArticle,
		},
	}

}

func startArticle(url string, method string) {

	newArticle = jianshu.NewArticle(url)
	if method == "get-author" || method == "1" {
		fmt.Println(newArticle.GetAuthor())
	}
	if method == "get-description" || method == "2" {
		fmt.Println(newArticle.GetDescription())
	}
	if method == "get-title" || method == "3" {
		fmt.Println(newArticle.GetTitle())
	}
	if method == "get-content" || method == "4" {
		fmt.Println(newArticle.GetContent())
	}
}
