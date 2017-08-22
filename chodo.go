package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/headzoo/surf.v1"
	"io/ioutil"
	"net/url"
	_ "runtime"
	str "strings"
	"sync"
)

var sent []string

func sendMessage(username string) bool {
	bow := surf.NewBrowser()
	err := bow.Open("https://" + username + ".sarahah.com/")
	if err != nil {
		return false
	}

	userid, _ := bow.Find("#RecipientId").Attr("value")
	text := `Hiiii! How is your job going? Are you enjoying your work or finding it boring? Speak your heart out!
The best part? You can do it anonymously. Go to Ambitionbox.cοm and rate your company now!`
	scriptCode := bow.Find("script").Text()
	if len(scriptCode) == 0 {
		return false
	}
	startI := str.Index(scriptCode, `type="hidden" value="`)
	endI := str.Index(scriptCode, `" />').attr('value')`)
	requestVerification := scriptCode[startI+len(`type="hidden" value="`) : endI]
	v := url.Values{}
	v.Set("userId", userid)
	v.Add("text", text)
	v.Add("captchaResponse", "")
	v.Add("__RequestVerificationToken", requestVerification)
	bow.PostForm("https://"+username+".sarahah.com/Messages/SendMessage", v)
	fmt.Println(username)
	return bow.Body() == "&#34;Done&#34;"
}

func main() {
	func() {
		raw, err := ioutil.ReadFile("./unames123.json")
		if err != nil {
			panic(err.Error())
		}

		var names []string
		json.Unmarshal(raw, &names)

		for _, name := range names {
			func() {
				if sendMessage(name) {
					sent = append(sent, name)
				}
			}()
		}
	}()
	func() {
		raw, err := ioutil.ReadFile("./unames.json")
		if err != nil {
			panic(err.Error())
		}

		var names []string
		json.Unmarshal(raw, &names)

		for _, name := range names {
			go func() {
				if sendMessage(name) {
					sent = append(sent, name)
				}
			}()
		}
	}()
	go func() {
		raw, err := ioutil.ReadFile("./ufirstnames.json")
		if err != nil {
			panic(err.Error())
		}

		var names []string
		json.Unmarshal(raw, &names)

		for _, name := range names {
			go func() {
				if sendMessage(name) {
					sent = append(sent, name)
				}
			}()
		}
	}()
	sentJson, _ := json.Marshal(sent)
	ioutil.WriteFile("sent.json", sentJson, 0644)
}
