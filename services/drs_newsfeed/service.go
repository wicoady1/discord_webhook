package drs_newsfeed

import (
	"fmt"

	"github.com/wicoady1/discord_webhook/util/gtrans_service"
)

//GenerateContent - Generate Content for Webhook
func GenerateContent() {
	param := gtrans_service.New(gtrans_service.ENGLISH_TRANSLATE, gtrans_service.JAPANESE_TRANSLATE, "testing")

	result, err := param.SendMessage()
	if err != nil {
		fmt.Println("%+v", err)
	}
	fmt.Printf(result)
}
