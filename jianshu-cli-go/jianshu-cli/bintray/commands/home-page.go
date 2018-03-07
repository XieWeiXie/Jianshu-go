package commands

import (
	"jianshu-go"

	"fmt"

	"github.com/urfave/cli"
)

var newHomePage *jianshu.HomePage

func GetHomePageFlag() []cli.Flag {
	return nil
}

func getHomePageFlag() []cli.Flag {
	return []cli.Flag{
		cli.IntFlag{
			Name:  "number",
			Usage: "the home page number to display",
		},
	}
}

func actionHomePage(c *cli.Context) {
	if c.NArg() < 1 {
		fmt.Println("execute command home-page method --number='*' <string>, should add one argument")
		return
	}
	if c.NArg() == 1 {
		startHomePage(c.Int("number"), c.Args()[0])
	}
}

func SubCMDHomePage() []cli.Command {
	return []cli.Command{
		{
			Name:   "method",
			Flags:  getHomePageFlag(),
			Usage:  "home-page method",
			Action: actionHomePage,
		},
	}

}

func startHomePage(number int, method string) {

	newHomePage = jianshu.NewHomePage(number, 0)
	if method == "get-home-page-passages" || method == "1" {
		fmt.Println(newHomePage.GetHomePagePassages())
	}
	if method == "get-new-list" || method == "2" {
		newHomePage.GetNewList()
	}
	if method == "get-hot-seven" || method == "3" {
		newHomePage.GetHotSeven()
	}
	if method == "get-hot-month" || method == "4" {
		newHomePage.GetHotMonth()
	}
	if method == "get-jianshu-school" || method == "5" {
		newHomePage.GetJianShuSchool()
	}
}
