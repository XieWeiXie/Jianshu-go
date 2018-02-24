package jianshu

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	rootUrl     = "https://www.jianshu.com/"
	rootUrlOnce = "https://www.jianshu.com"
)

func MakeCompleteUrl(element string) string {
	if element == "" {
		panic("element should not be nil.")
	}
	if strings.HasPrefix(element, "/") {
		return rootUrlOnce + element
	}
	return rootUrl + element
}

func SplitString(element string) string {
	newElement := strings.SplitAfter(element, "-")
	return newElement[1]
}

func StringToInt(element string) int {
	newInt, _ := strconv.Atoi(element)
	return newInt
}

func Pages(allPassage int) int {
	return allPassage/9 + 1
}

func MakeCompletePages(page int, user *User) string {
	return fmt.Sprintf(user.Link+"?order_by=shared_at&page=%s", strconv.Itoa(page))
}

func StringSpace(element string) string {
	return strings.TrimSpace(element)
}

func GetUuidFromLink(user *User) string {
	if user.Link == "" {
		return "None"
	}
	urlList := strings.Split(user.Link, "/")
	return urlList[4]
}

func LikedNotesUrl(user *User) string {
	id := GetUuidFromLink(user)
	if id == "None" {
		panic("id should not be nil.")
	}
	return rootUrl + "users/" + id + "/liked_notes?_pjax=%23list-container"
}

func SubscriptionUrl(user *User) string {
	id := GetUuidFromLink(user)
	if id == "None" {
		panic("id should not be nil.")
	}
	return rootUrl + "users/" + id + "/subscriptions?_pjax=%23list-container"
}

func StringGetInt(value string) int {
	tempInt := strings.SplitAfter(value, " ")
	return StringToInt(tempInt[1])
}

func StringHandle(value string) (string, string) {
	newValue := StringSpace(strings.Replace(value, "\n", "", -1))
	newValueList := strings.Split(newValue, " ")
	fmt.Println(newValueList)
}

func RegexpNumber(value string) (int, int) {

}
