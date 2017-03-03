package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Attachment struct {
	Color  string  `json:"color"`
	Fields []Field `json:"fields"`
}

type Payload struct {
	Text        string       `json:"text"`
	Username    string       `json:"username"`
	Channel     string       `json:"channel"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

func PostToSlack(text string, attributes map[string]string, slackConfig SlackConfig) string {
	if slackConfig.Url == "" {
		return "pass"
	}

	fields := []Field{}
	for key, value := range attributes {
		fields = append(fields, Field{key, value, true})
	}

	params, _ := json.Marshal(Payload{
		text,
		slackConfig.Username,
		slackConfig.Channel,
		[]Attachment{Attachment{"danger", fields}}})

	resp, _ := http.PostForm(
		slackConfig.Url,
		url.Values{"payload": {string(params)}},
	)

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return string(body)
}
