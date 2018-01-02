package webhook

func InitWebHookModule() *WebhookModule {
	discordURL := "<place your webhook API URL here>"

	config := Config{
		URL: discordURL,
	}

	module := WebhookModule{
		Config: &config,
	}

	return &module
}
