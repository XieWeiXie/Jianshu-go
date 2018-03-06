package bintray

import (
	"jianshu-go/jianshu-cli-go/jianshu-cli/bintray/commands"

	"jianshu-go/jianshu-cli-go/jianshu-cli/utils"

	"github.com/urfave/cli"
)

func GetCommands() []cli.Command {
	return []cli.Command{
		{
			Name:        "article",
			Usage:       utils.GetUsage("article"),
			UsageText:   utils.GetUsageText("article"),
			HelpName:    utils.GetHelp("article"),
			Flags:       commands.GetArticleFlag(),
			Subcommands: commands.SubCMDArticle(),
		},
		{
			Name:        "home-page",
			Usage:       utils.GetUsage("home-page"),
			UsageText:   utils.GetUsageText("home-page"),
			HelpName:    utils.GetHelp("home-page"),
			Flags:       commands.GetHomePageFlag(),
			Subcommands: commands.SubCMDHomePage(),
		},
		{
			Name:        "recommend",
			Usage:       utils.GetUsage("recommend"),
			UsageText:   utils.GetUsageText("recommend"),
			HelpName:    utils.GetHelp("recommend"),
			Flags:       commands.GetHomePageRecommendFlag(),
			Subcommands: commands.SubCMDHomePageRecommend(),
		},
		{
			Name:        "topic",
			Usage:       utils.GetUsageText("topic"),
			UsageText:   utils.GetUsageText("topic"),
			HelpName:    utils.GetHelp("topic"),
			Flags:       commands.GetTopicFlag(),
			Subcommands: commands.SubCMDTopic(),
		},
		{
			Name:        "user",
			Usage:       utils.GetUsage("user"),
			UsageText:   utils.GetUsageText("user"),
			HelpName:    utils.GetHelp("user"),
			Flags:       commands.GetUserFlag(),
			Subcommands: commands.SubCMDUser(),
		},
		{
			Name:        "publication",
			Flags:       commands.GetPublicationFlag(),
			Subcommands: commands.SubCMDPublication(),
		},
	}
}
