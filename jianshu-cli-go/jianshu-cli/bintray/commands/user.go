package commands

import (
	"fmt"

	"jianshu-go"

	"github.com/urfave/cli"
)

func GetUserFlag() []cli.Flag {
	return nil
}

var newUser *jianshu.User

func actionUserInit(cli *cli.Context) {
	if cli.NArg() < 1 {
		fmt.Println("execute command user init should add one link")
		return
	}
	if cli.NArg() == 1 {
		start(cli.Args()[0])
	}
}

func actionUserMethod(cli *cli.Context) {
	if cli.NArg() < 1 {
		fmt.Println("execute command user method should add one argument")
		return
	}
	if cli.NArg() == 1 {
		if cli.Args()[0] == "get-user-id" {
			getUserID()
		}

	}
}

func SubCMDUser() []cli.Command {
	return []cli.Command{
		{
			Name:   "init",
			Usage:  "start init newUser",
			Action: actionUserInit,
		},
		{
			Name:   "method",
			Usage:  "get NewUser method",
			Action: actionUserMethod,
		},
	}
}

func start(link string) {
	newUser = jianshu.NewUser(link, "", "")
	fmt.Println(newUser.GetUserID())
}

func getUserID() {
	fmt.Println(newUser.GetUserID())
}
