package webhook

/*
content	string	the message contents (up to 2000 characters)	one of content, file, embeds
username	string	override the default username of the webhook	false
avatar_url	string	override the default avatar of the webhook	false
tts	bool	true if this is a TTS message	false
file	file contents	the contents of the file being sent	one of content, file, embeds
embeds	array of embed objects	embedded rich content	one of content, file, embeds
*/
type WebhookModule struct {
	Config *Config
}

type Config struct {
	URL string
}

type WebhookContent struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds"`
}

/*
title	string	title of embed
type	string	type of embed (always "rich" for webhook embeds)
description	string	description of embed
url	string	url of embed
timestamp	ISO8601 timestamp	timestamp of embed content --> 2017-11-28T13:55:36+00:00
color	integer	color code of the embed
footer	embed footer object	footer information
image	embed image object	image information
thumbnail	embed thumbnail object	thumbnail information
video	embed video object	video information
provider	embed provider object	provider information
author	embed author object	author information
fields	array of embed field objects	fields information
*/

type Embed struct {
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Timestamp   string `json:"timestamp"` //yyyy-mm-dd
	Color       int    `json:"color"`
}
