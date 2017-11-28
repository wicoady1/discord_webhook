package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/wicoady1/discord-webhook/database"
)

type Action interface {
	SendMessage(WebhookContent)
	GenerateContent() WebhookContent
}

//SendMessage - Send message to Webhook, according to provided content
func (wm *WebhookModule) SendMessage(content WebhookContent) {
	url := wm.Config.URL

	jsonStr, err := json.Marshal(content)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

//GenerateContent - Generate Content for Webhook
func (wm *WebhookModule) GenerateContent() WebhookContent {
	articleList, err := database.GetPendingArticleList()
	if err != nil {
		log.Println("[GenerateContent]", err)
		return WebhookContent{}
	}

	contentEmbed := []Embed{}
	for _, v := range articleList {
		embed := Embed{}

		embed.Title = v.Title
		embed.Description = fmt.Sprintf("Current Status: %s\nOriginally written by %s", v.Status, v.AuthorName)
		embed.Type = "rich"

		embed.Color = 16724787 //status --> Draft
		if v.Status == "pending" {
			embed.Color = 16777011 //status --> Pending
		}

		contentEmbed = append(contentEmbed, embed)
	}

	return WebhookContent{
		Content: "@here Unpublished Article List",
		Embeds:  contentEmbed,
	}
}
