package gtrans_service

const GOOGLE_TRANSLATE_API = "https://translate.googleapis.com/translate_a/single"

/*
Translate Param:
SourceLanguage: source language
TranslateLanguage: translated to language
Query: text to be translated
*/

type TranslateParam struct {
	SourceLanguage    string `url:"sl"`
	TranslateLanguage string `url:"tl"`
	Query             string `url:"q"`
	Client            string `url:"client"`
	Dt                string `url:"dt"`
}

const ENGLISH_TRANSLATE string = "en"
const JAPANESE_TRANSLATE string = "ja"
const DEFAULT_CLIENT string = "dtx"
const DEFAULT_DT string = "t"

type Result interface{}
