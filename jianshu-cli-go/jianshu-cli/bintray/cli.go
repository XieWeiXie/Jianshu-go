package bintray

import "github.com/urfave/cli"

func GetCommands() []cli.Command {
	return []cli.Command{
		{
			Name: "article",
		},
		{
			Name: "home-page",
		},
		{
			Name: "recommend",
		},
		{
			Name: "topic",
		},
		{
			Name: "user",
		},
		{
			Name: "publication",
		},
	}
}
