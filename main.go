package main

import (
	"log"
	"net/http"

	"github.com/robfig/cron"
	"github.com/wicoady1/discord_webhook/services/drs_newsfeed"
	"github.com/wicoady1/discord_webhook/util/database"
	"github.com/wicoady1/discord_webhook/webhook"
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

	err := cronTask.AddFunc("0/10 * * * * *", func() {
		/*
			content := module.GenerateContent()
			module.SendMessage(content)
		*/
		drs_newsfeed.GenerateContent()
	})
	if err != nil {
		log.Println("[CRON Task]", err)
		return
	}

	cronTask.Start()
	log.Println("[Webhook Bot] Cron task is running!")

	http.ListenAndServe(":8000", nil) //dummy to enforce program to keep running
}
