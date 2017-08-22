package main

import (
	"fmt"
	"gopkg.in/headzoo/surf.v1"
	"net/url"
	str "strings"
)

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("https://ambitionbox.sarahah.com/")
	if err != nil {
		panic(err)
	}

	userid, _ := bow.Find("#RecipientId").Attr("value")
	text := "Test from Go"
	scriptCode := bow.Find("script").Text()
	startI := str.Index(scriptCode, `type="hidden" value="`)
	endI := str.Index(scriptCode, `" />').attr('value')`)
	requestVerification := scriptCode[startI+len(`type="hidden" value="`) : endI]
	v := url.Values{}
	v.Set("userId", userid)
	v.Add("text", text)
	v.Add("captchaResponse", "")
	v.Add("__RequestVerificationToken", requestVerification)
	bow.PostForm("https://ambitionbox.sarahah.com/Messages/SendMessage", v)
	fmt.Println(bow.Body())

}
