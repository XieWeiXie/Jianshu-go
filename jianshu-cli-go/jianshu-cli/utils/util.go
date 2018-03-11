package utils

import (
	"fmt"
	"strings"
)

const (
	AppName    = "JianShu"
	AppUsage   = "An Application of JianShu API."
	AppVersion = `

    ___       ___       ___       ___       ___       ___       ___   
   /\  \     /\  \     /\  \     /\__\     /\  \     /\__\     /\__\  
  _\:\  \   _\:\  \   /::\  \   /:| _|_   /::\  \   /:/__/_   /:/ _/_ 
 /\/::\__\ /\/::\__\ /::\:\__\ /::|/\__\ /\:\:\__\ /::\/\__\ /:/_/\__\
 \::/\/__/ \::/\/__/ \/\::/  / \/|::/  / \:\:\/__/ \/\::/  / \:\/:/  /
  \/__/     \:\__\     /:/  /    |:/  /   \::/  /    /:/  /   \::/  / 
             \/__/     \/__/     \/__/     \/__/     \/__/     \/__/  v1.0.0


`
)

// help
const (
	articleHelp      = "get jianshu article"
	articleUsageText = "get jianshu article"
	articleUsage     = "get jianshu article by cli"
)

const (
	homePageHelp      = "get home page passage"
	homePageUsageText = "get home page passage"
	homePageUsage     = "get jianshu home page passage by cli"
)

const (
	homePageRecommendHelp      = "get home page recommend"
	homePageRecommendUsageText = "get home page recommend"
	homePageRecommendUsage     = "get home page recommend by cli"
)

const (
	homePageTopicHelp      = "get home page topic"
	homePageTopicUsageText = "get home page topic"
	homePageTopicUsage     = "get home page topic by cli"
)

const (
	publicationHelp      = "get publication"
	publicationUsageText = "get publication"
	publicationUsage     = "get publication by cli"
)

const (
	userHelp      = "get user"
	userUsageText = "get user"
	userUsage     = "get user by cli"
)

func GetUsage(key string) string {
	var result string
	if strings.Contains(key, "article") {
		result = articleUsage
	}
	if strings.Contains(key, "home-page") {
		result = homePageUsage
	}
	if strings.Contains(key, "recommend") {
		result = homePageRecommendUsage
	}
	if strings.Contains(key, "topic") {
		result = homePageTopicUsage
	}
	if strings.Contains(key, "publication") {
		result = publicationUsage
	}
	if strings.Contains(key, "user") {
		result = userUsage
	}
	return result
}

func GetUsageText(key string) string {
	var result string
	if strings.Contains(key, "article") {
		result = articleUsageText
	}
	if strings.Contains(key, "home-page") {
		result = homePageUsageText
	}
	if strings.Contains(key, "recommend") {
		result = homePageRecommendUsageText
	}
	if strings.Contains(key, "topic") {
		result = homePageTopicUsageText
	}
	if strings.Contains(key, "publication") {
		result = publicationUsageText
	}
	if strings.Contains(key, "user") {
		result = userUsageText
	}
	return result
}

func GetHelp(key string) string {
	var result string
	if strings.Contains(key, "article") {
		result = articleHelp
	}
	if strings.Contains(key, "home-page") {
		result = homePageHelp
	}
	if strings.Contains(key, "recommend") {
		result = homePageRecommendHelp
	}
	if strings.Contains(key, "topic") {
		result = homePageTopicHelp
	}
	if strings.Contains(key, "publication") {
		result = publicationHelp
	}
	if strings.Contains(key, "user") {
		result = userHelp
	}
	return result
}

func HelpExample() string {
	fmt.Println("谢小路：jianshu-link: https://www.jianshu.com/u/58f0817209aa")
	var link = "https://www.jianshu.com/u/58f0817209aa"
	return fmt.Sprintf("Jianshu user --link=%s "+"method 1", link)
}
