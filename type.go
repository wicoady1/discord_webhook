package main

/*
content	string	the message contents (up to 2000 characters)	one of content, file, embeds
username	string	override the default username of the webhook	false
avatar_url	string	override the default avatar of the webhook	false
tts	bool	true if this is a TTS message	false
file	file contents	the contents of the file being sent	one of content, file, embeds
embeds	array of embed objects	embedded rich content	one of content, file, embeds
*/
type WebhookContent struct {
	Content string `json:"content"`
}
