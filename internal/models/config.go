package models

type Bot struct {
	Token    string   `json:"token"`
	Channels []string `json:"channels"`
	Trigger  string   `json:"trigger"`
	Interval string   `json:"interval"`
	Mentions []string `json:"mentions"`
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

type Config struct {
	Bot       Bot
	Scrape    Scrape
	Translate Translate
}
