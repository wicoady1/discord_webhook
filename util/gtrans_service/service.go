package gtrans_service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
)

//New - Initialize TranslateParam
func New(sourceLang string, destinationLang string, query string) TranslateParam {
	return TranslateParam{
		SourceLanguage:    sourceLang,
		TranslateLanguage: destinationLang,
		Query:             query,
		Dt:                DEFAULT_DT,
		Client:            DEFAULT_CLIENT,
	}
}

//SendMessage - Send message to Webhook, according to provided content
func (tp *TranslateParam) SendMessage() (string, error) {
	url := tp.generateURL()
	fmt.Printf("[REQ] \n%+v\n", url)

	client := &http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println("[RES] Status:", resp.Status)
	fmt.Println("[RES] Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("[RES] Body:", string(body))

	result, err := tp.extractTranslationResult(body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return result, nil
}

//generateURL - Generate URL for call
func (tp *TranslateParam) generateURL() string {
	v, _ := query.Values(tp)

	return fmt.Sprintf("%s?%s", GOOGLE_TRANSLATE_API, v.Encode())
}

//extractTranslationResult - Extract Translation Result
func (tp *TranslateParam) extractTranslationResult(raw []byte) (string, error) {
	var data interface{}

	err := json.Unmarshal(raw, &data)
	if err != nil {
		return "", err
	}

	//brutally extract translation result from nested array response
	depthCounter := 0
	currentData := data
	for depthCounter < 3 {
		depthCounter++

		if reflect.ValueOf(currentData).Kind() == reflect.Slice {
			d := reflect.ValueOf(currentData)
			currentData = d.Index(0).Interface()
		}
	}

	//result is not located there error
	if reflect.ValueOf(currentData).Kind() != reflect.String {
		return "", fmt.Errorf("failed to parse to string: %s", currentData)
	}

	return currentData.(string), nil
}
