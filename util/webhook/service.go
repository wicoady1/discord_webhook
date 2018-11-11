package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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

	client := &http.Client{
		Timeout: time.Second * 1,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

//GenerateContent - Generate Content for Webhook
func (wm *WebhookModule) GenerateContent(title string, description string) WebhookContent {
	/*
		articleList, err := database.GetPendingArticleList()
		if err != nil {
			log.Println("[GenerateContent]", err)
			return WebhookContent{}
		}
	*/

	contentEmbed := []Embed{}
	embed := Embed{}

	embed.Title = title
	embed.Description = description
	embed.Type = "rich"
	embed.Color = 16724787 //status --> Draft

	contentEmbed = append(contentEmbed, embed)

	return WebhookContent{
		Content: "@here Testing!",
		Embeds:  contentEmbed,
	}
}
