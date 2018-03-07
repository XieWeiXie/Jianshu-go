package commands

import (
	"jianshu-go"

	"fmt"

	"github.com/urfave/cli"
)

var newTopic *jianshu.Topic

func GetTopicFlag() []cli.Flag {
	return nil
}

func getTopicFlag() []cli.Flag {
	return []cli.Flag{
		cli.IntFlag{
			Name:  "number",
			Usage: "the number of topic to display",
		},
	}
}

func actionTopic(c *cli.Context) {
	if c.NArg() < 1 {
		fmt.Println("execute command topic method --number='*' <string>, should add one argument")
		return
	}
	if c.NArg() == 1 {
		startTopic(c.Int("number"), c.Args()[0])
	}
}

func SubCMDTopic() []cli.Command {
	return []cli.Command{
		{
			Name:   "method",
			Usage:  "the method of topic",
			Flags:  getTopicFlag(),
			Action: actionTopic,
		},
	}
}

func startTopic(number int, method string) {
	newTopic = jianshu.NewTopic(number)
	if method == "get-topic-collection-recommend" || method == "1" {
		fmt.Println(newTopic.GetTopicCollectionRecommend())
	}
	if method == "get-topic-collection-hot" || method == "2" {
		fmt.Println(newTopic.GetTopicCollectionHot())
	}
	if method == "get-topic-collection-city" || method == "3" {
		fmt.Println(newTopic.GetTopicCollectionCity())
	}
	if method == "get-topic-collection-schoolyard" || method == "4" {
		fmt.Println(newTopic.GetTopicCollectionSchoolyard())
	}
}
