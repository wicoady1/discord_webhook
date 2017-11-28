package webhook

import (
	"io/ioutil"
	"log"
	"strings"
)

func InitWebHookModule() *WebhookModule {
	module, err := ParseConfigFile()
	if err != nil {
		panic(err)
	}

	return module
}

func ParseConfigFile() (*WebhookModule, error) {
	module := &WebhookModule{}

	configFile, err := ioutil.ReadFile("config/config.ini")
	if err != nil {
		log.Println("[Config File Err]", err)
		return module, err
	}

	var configAll Config

	configs := strings.Split(string(configFile), "\n")
	for _, v := range configs {
		each := strings.Split(v, "|")
		switch each[0] {
		case "database":
			configAll.Database = each[1]
		case "webhook":
			configAll.URL = each[1]
		}
	}

	module.Config = &configAll

	return module, nil
}
