package commands

import (
	"fmt"

	"github.com/wuxiaoxiaoshen/jianshu-go"

	"github.com/wuxiaoxiaoshen/jianshu-go/jianshu-cli-go/jianshu-cli/utils"

	"github.com/urfave/cli"
)

func GetUserFlag() []cli.Flag {
	return nil
}

var newUser *jianshu.User

func getFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "url",
			Usage: "jianshu personal url",
		},
	}
}
func actionUserMethod(cli *cli.Context) {
	if cli.NArg() < 1 {
		fmt.Println("execute command user method --link='****' <string> , should add one argument")
		return
	}
	if cli.NArg() == 1 {
		startUser(cli.String("url"), cli.Args()[0])
	}
}

func SubCMDUser() []cli.Command {
	return []cli.Command{
		{
			Name:   "method",
			Usage:  "get NewUser method",
			Flags:  getFlag(),
			Action: actionUserMethod,
		},
	}
}

func startUser(link string, method string) {
	if link == "" {
		fmt.Println("Please add jianshu link.")
		link = utils.HelpExample()
		fmt.Printf("link: %s \n", link)
		return
	}
	if method == "" {
		fmt.Println("please add user method.")
		return
	}
	newUser = jianshu.NewUser(link, "", "")
	if method == "get-user-id" || method == "1" {
		fmt.Println(newUser.GetUserID())
	}
	if method == "get-user-gender" || method == "2" {
		fmt.Println(newUser.GetUserGender())
	}
	if method == "get-user-link" || method == "3" {
		fmt.Println(newUser.GetUserLink())
	}
	if method == "get-follow-number" || method == "4" {
		fmt.Println(newUser.GetFollowNumber())
	}
	if method == "get-follower-number" || method == "5" {
		fmt.Println(newUser.GetFollowerNumber())
	}
	if method == "get-passage-number" || method == "6" {
		fmt.Println(newUser.GetPassageNumber())
	}
	if method == "get-write-number" || method == "7" {
		fmt.Println(newUser.GetWriteNumber())
	}
	if method == "get-like-number" || method == "8" {
		fmt.Println(newUser.GetLikeNumber())
	}
	if method == "get-home-page-passage" || method == "9" {
		fmt.Println(newUser.GetHomepagePassage())
	}
	if method == "get-personal-detail" || method == "10" {
		fmt.Println(newUser.GetPersonalDetail())
	}
	if method == "get-twitter-info" || method == "11" {
		fmt.Println(newUser.GetTwitterInfo())
	}
	if method == "get-liked-notes" || method == "12" {
		fmt.Println(newUser.GetLikedNotes())
	}
	if method == "get-subscription" || method == "13" {
		fmt.Println(newUser.GetSubscription())
	}
	if method == "get-latest-activate" || method == "14" {
		fmt.Println(newUser.GetLatestActive())
	}
	if method == "get-latest-comment" || method == "15" {
		fmt.Println(newUser.GetLatestCommented())
	}
	if method == "get-hot-pages" || method == "16" {
		fmt.Println(newUser.GetHotPassage())
	}

}
