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
	fmt.Println("URL:>", url)

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
	return WebhookContent{
		Content: "Test Webhook",
		Embeds: []Embed{
			Embed{
				Title:       "Testing",
				Type:        "rich",
				Description: "testing by ken(pai)",
				URL:         "https://google.com",
				Timestamp:   time.Now().Format("2006-01-02T15:04:05+0700"),
				Color:       14177041,
			},
		},
	}
}
