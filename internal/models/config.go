package models

type Bot struct {
	Token           string   `json:"token"`
	Channels        []string `json:"channels"`
	Trigger         string   `json:"trigger"`
	Interval        string   `json:"interval"`
	Mentions        []string `json:"mentions"`
	PrivilegedUsers []string `json:"privileged_users"`
}

type Scrape struct {
	Instances []string `json:"instances"`
	Interval  string   `json:"interval"`
	Usernames []string `json:"usernames"`
	Redirect  string   `json:"redirect"`
}

type DeepL struct {
	Url string `json:"url"`
	Key string `json:"key"`
}

type Translate struct {
	DeepL DeepL `json:"deepl"`
}

type OpenAI struct {
	Model  string `json:"model"`
	Url    string `json:"url"`
	Key    string `json:"key"`
	System string `json:"system"`
}

type Chat struct {
	Public bool   `json:"public"`
	OpenAI OpenAI `json:"openai"`
}

type Config struct {
	Bot       Bot
	Scrape    Scrape
	Translate Translate
	Chat      Chat
}
