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
func (wm *WebhookModule) GenerateContent() WebhookContent {
	/*
		articleList, err := database.GetPendingArticleList()
		if err != nil {
			log.Println("[GenerateContent]", err)
			return WebhookContent{}
		}
	*/

	fmt.Printf("masuk pak eko")
	contentEmbed := []Embed{}
	embed := Embed{}

	embed.Title = "Test Discord DRS"
	embed.Description = fmt.Sprintf("Test MakLO! udah connect gak?")
	embed.Type = "rich"
	embed.Color = 16724787 //status --> Draft

	contentEmbed = append(contentEmbed, embed)

	return WebhookContent{
		Content: "@here MAKLO!",
		Embeds:  contentEmbed,
	}
}
