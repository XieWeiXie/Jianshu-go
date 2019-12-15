package commands

import (
	"github.com/wuxiaoxiaoshen/jianshu-go"

	"fmt"

	"github.com/urfave/cli"
)

var newPublication *jianshu.Publication

func GetPublicationFlag() []cli.Flag {
	return nil
}

func getPublicationFlag() []cli.Flag {
	return nil
}

func actionPublication(c *cli.Context) {
	if c.NArg() < 1 {
		fmt.Println("execute command publication method <string>, should add one argument")
		return
	}
	if c.NArg() == 1 {
		startPublication(c.Args()[0])
	}
}

func SubCMDPublication() []cli.Command {
	return []cli.Command{
		{
			Name:   "method",
			Usage:  "get publication method",
			Flags:  getPublicationFlag(),
			Action: actionPublication,
		},
	}
}

func startPublication(method string) {
	newPublication = jianshu.NewPublication()
	if method == "get-publicized-book" || method == "1" {
		newPublication.GetPublicizedBook()
	}
	if method == "get-novel-books" || method == "2" {
		newPublication.GetNovelBooks()
	}
	if method == "get-it-job-market" || method == "3" {
		newPublication.GetITAndJobMarket()
	}
	if method == "get-culture-history" || method == "4" {
		newPublication.GetCultureAndHistory()
	}
	if method == "get-monthly-magazine" || method == "5" {
		newPublication.GetMonthlyMagazine()
	}
}
