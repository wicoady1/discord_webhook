package main

import "github.com/wicoady1/discord-webhook/webhook"

func main() {
	module := webhook.InitWebHookModule()

	content := module.GenerateContent()
	module.SendMessage(content)
}
