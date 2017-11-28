package main

import (
	"log"
	"net/http"

	"github.com/robfig/cron"
	"github.com/wicoady1/discord-webhook/database"
	"github.com/wicoady1/discord-webhook/webhook"
)

var cronTask *cron.Cron

func main() {
	//initialize DB
	module := webhook.InitWebHookModule()
	database.Init(module.Config.Database)

	//initialize Cron
	if cronTask == nil {
		cronTask = cron.New()
	} else {
		cronTask.Stop()
	}

	err := cronTask.AddFunc("0 0 0 * * *", func() {
		content := module.GenerateContent()
		module.SendMessage(content)
	})
	if err != nil {
		log.Println("[CRON Task]", err)
		return
	}

	cronTask.Start()
	log.Println("[AG.ID Webhook Bot] Cron task is running!")

	http.ListenAndServe(":8080", nil) //dummy to enforce program to keep running
}
