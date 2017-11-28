package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://discordapp.com/api/webhooks/385086724163764224/_OF9O1EuhlAWI8dpD8cPV9pQNlArRkf4oCQsFO4FkoA1TGpKEzEwpmwx8aleuVw5b8hf"
	fmt.Println("URL:>", url)

	content := WebhookContent{
		Content: "Test Bosq!",
	}

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
