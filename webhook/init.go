package webhook

func InitWebHookModule() *WebhookModule {
	discordURL := "https://discordapp.com/api/webhooks/385086724163764224/_OF9O1EuhlAWI8dpD8cPV9pQNlArRkf4oCQsFO4FkoA1TGpKEzEwpmwx8aleuVw5b8hf"

	config := Config{
		URL: discordURL,
	}

	module := WebhookModule{
		Config: &config,
	}

	return &module
}
