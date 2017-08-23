package main

import (
	"encoding/json"
	"fmt"
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"
	"io/ioutil"
	"net/url"
	_ "runtime"
	str "strings"
)

var sent []string

var bow *browser.Browser

func logIn() {
	bow = surf.NewBrowser()
	bow.Open("https://sarahah.com/Account/Login")
	form, _ := bow.Form("form")
	form.Input("Email", "ambitionbox")

	form.Submit()
	fmt.Println(bow.Url())
}

func sendMessage(username string) bool {
	err := bow.Open("https://" + username + ".sarahah.com/")
	if err != nil || len(bow.Body()) < 7000 {
		return false
	}
	// cookie := `.AspNetCore.Antiforgery.w5W7x28NAIs=CfDJ8ATf4Su878dOnC1xwa5QU5tOYQXGJea6T6emsXmkyhZ6Wa-irWebbKywSvk52MUZBnPpi1uipEXjYnFLqEzDirAs3gk8H40AKcZCMYR5zzfSLclQOMAmmmi39znXOFcPsUf_D7njULUqBnX2FCsXbFE; ai_user=0aMLg|2017-08-22T16:55:14.990Z; .AspNetCore.Identity.Application=CfDJ8ATf4Su878dOnC1xwa5QU5uZStJHqzq7sL2HWaJjmU0N3qJbvCBFEMLDDQEaL4BqRQip-KPOzabMJGVttNKjxSUbA6g1Oh4C4hUZqeJHtW-WckhDihVPTM5b0lI-uaZwkDKEgMXKJkbaRLntmxFCgwklIQKkmxQ3hDEGX-DqRt2F70HG0YnqYxXdYqnGWXwfKK6GrgD00ll5az7H7bmk3z7hSq0GZoKYzpPHMwJefomeFGDFM9XrOloBvxAbyFJCFMjqAP1uA9CPlAB4Ai1mL7dr9hf_UH6W-qmvNG36QSa5PQL0Hgj9KDJJl-PO7QXupmNgIR6EgZCZWWQrqDa53P9bHG6ApujU7WTcTuXbFAPw60eQKJgOkpVEo1mfo68LkXSVJvXAbPMUTTKqBTKpRj1qxPmvPIuVQ61Geee8adh41MdfQYtAcUeUpBho05sonZCD5SfFoUrWLytOTHyuQnXZtJ8dMzCC1rOWQ7RGa23zJsFk3C3hktMFmxQeq6dXFXBbtTpTww54hkkyLj4MJE8; ai_session=KzOR8|1503471846954.95|1503471846954.95`
	// bow.AddRequestHeader("Cookie", cookie)

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
	logIn()
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
	raw, err = ioutil.ReadFile("./unames.json")
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(raw, &names)

	for _, name := range names {
		func() {
			if sendMessage(name) {
				sent = append(sent, name)
			}
		}()
	}
	raw, err = ioutil.ReadFile("./ufirstnames.json")
	if err != nil {
		panic(err.Error())
	}

	json.Unmarshal(raw, &names)

	for _, name := range names {
		func() {
			if sendMessage(name) {
				sent = append(sent, name)
			}
		}()
	}
	sentJson, _ := json.Marshal(sent)
	ioutil.WriteFile("sent.json", sentJson, 0644)
}
