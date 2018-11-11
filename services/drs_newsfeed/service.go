package drs_newsfeed

import (
	"fmt"

	"github.com/wicoady1/discord_webhook/util/gtrans_service"
	"github.com/wicoady1/discord_webhook/util/webhook"
)

//GenerateContent - Generate Content for Webhook
func GenerateContent(wm *webhook.WebhookModule) {
	testString := "DanceRush Stardom is already available!"

	param := gtrans_service.New(gtrans_service.ENGLISH_TRANSLATE, gtrans_service.JAPANESE_TRANSLATE, testString)

	result, err := param.SendMessage()
	if err != nil {
		fmt.Println("%+v", err)
	}

	content := wm.GenerateContent("Welcumm", fmt.Sprintf("ENG: %s\nJPN: %s", testString, result))
	wm.SendMessage(content)
}
